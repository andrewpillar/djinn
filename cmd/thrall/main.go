package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/andrewpillar/cli"

	"github.com/andrewpillar/thrall/collector"
	"github.com/andrewpillar/thrall/config"
	"github.com/andrewpillar/thrall/driver"
	"github.com/andrewpillar/thrall/placer"
	"github.com/andrewpillar/thrall/runner"
)

var setupStage = "setup"

func initializeSSH(manifest config.Manifest) runner.Driver {
	timeout, err := strconv.ParseInt(os.Getenv("THRALL_SSH_TIMEOUT"), 10, 64)

	if err != nil {
		timeout = 10
	}

	return &driver.SSH{
		Writer:   os.Stdout,
		Address:  manifest.Driver.Address,
		Username: manifest.Driver.Username,
		KeyFile:  os.Getenv("THRALL_SSH_KEY"),
		Timeout:  time.Duration(time.Second * time.Duration(timeout)),
	}
}

func initializeQEMU(manifest config.Manifest) runner.Driver {
	hostfwd := os.Getenv("THRALL_QEMU_HOSTFWD")

	if hostfwd == "" {
		hostfwd = "127.0.0.1:2222"
	}

	timeout, err := strconv.ParseInt(os.Getenv("THRALL_SSH_TIMEOUT"), 10, 64)

	if err != nil {
		timeout = 10
	}

	ssh := &driver.SSH{
		Writer:   ioutil.Discard,
		Address:  hostfwd,
		Username: os.Getenv("THRALL_QEMU_USERNAME"),
		KeyFile:  os.Getenv("THRALL_SSH_KEY"),
		Timeout:  time.Duration(time.Second * time.Duration(timeout)),
	}

	driver.QemuDir = os.Getenv("THRALL_QEMU_DIR")

	cpus, err := strconv.ParseInt(os.Getenv("THRALL_QEMU_CPUS"), 10, 64)

	if err == nil {
		driver.QemuCPUs = cpus
	}

	memory, err := strconv.ParseInt(os.Getenv("THRALL_QEMU_MEMORY"), 10, 64)

	if err == nil {
		driver.QemuMemory = memory
	}

	arch := manifest.Driver.Arch

	if arch == "" {
		arch = "x86_64"
	}

	return &driver.QEMU{
		Writer:  os.Stdout,
		SSH:     ssh,
		Image:   manifest.Driver.Image,
		Arch:    arch,
		HostFwd: hostfwd,
	}
}

func mainCommand(c cli.Command) {
	f, err := os.Open(c.Flags.GetString("manifest"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	manifest, err := config.DecodeManifest(f)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	if err := manifest.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	sigs := make(chan os.Signal)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGKILL)

	r := runner.NewRunner(
		os.Stdout,
		manifest.Env,
		manifest.Objects,
		placer.NewFileSystem(c.Flags.GetString("objects")),
		collector.NewFileSystem(c.Flags.GetString("artifacts")),
		sigs,
	)

	setup := runner.NewStage(setupStage, false)

	for i, src := range manifest.Sources {
		name := fmt.Sprintf("clone.%d", i + 1)

		commands := []string{
			"git clone " + src.URL + " " + src.Dir,
			"cd " + src.Dir,
			"git checkout -q " + src.Ref,
		}

		if src.Dir != "" {
			commands = append([]string{"mkdir -p " + src.Dir}, commands...)
		}

		depends := []string{"create driver"}

		if i > 0 {
			depends = append(depends, fmt.Sprintf("clone.%d", i))
		}

		setup.Add(runner.NewJob(os.Stdout, name, commands, depends, []config.Passthrough{}))
	}

	r.Add(setup)

	for _, name := range manifest.Stages {
		canFail := false

		for _, search := range manifest.AllowFailures {
			if name == search {
				canFail = true
				break
			}
		}

		r.Add(runner.NewStage(name, canFail))
	}

	for _, j := range manifest.Jobs {
		if _, ok := r.Stages[j.Stage]; !ok {
			fmt.Fprintf(r.Writer, "warning: unknown stage %s\n", j.Stage)
		}
	}

	for _, s := range r.Stages {
		jobId := 1

		for _, j := range manifest.Jobs {
			if s.Name != j.Stage {
				continue
			}

			if j.Name == "" {
				j.Name = fmt.Sprintf("%s.%d", j.Stage, jobId)
				jobId++
			}

			s.Add(runner.NewJob(os.Stdout, j.Name, j.Commands, j.Depends, j.Artifacts))
		}
	}

	var d runner.Driver

	switch manifest.Driver.Type {
		case "docker":
			d = driver.NewDocker(manifest.Driver.Image, manifest.Driver.Workspace)
		case "qemu":
			d = initializeQEMU(manifest)
		case "ssh":
			d = initializeSSH(manifest)
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown driver %s\n", os.Args[0], manifest.Driver.Type)
			os.Exit(1)
	}

	stages := c.Flags.GetAll("stage")

	if len(stages) > 0 {
		remove := make([]string, 0, len(r.Stages))

		for runnerStage := range r.Stages {
			keep := false

			for _, flag := range stages {
				if runnerStage == flag.Value || runnerStage == setupStage {
					keep = true
				}
			}

			if !keep {
				remove = append(remove, runnerStage)
			}
		}

		r.Remove(remove...)
	}

	if err := r.Run(d); err != nil {
		os.Exit(1)
	}
}

func main() {
	c := cli.New()

	c.AddFlag(&cli.Flag{
		Name:      "help",
		Long:      "--help",
		Exclusive: true,
		Handler:   func(f cli.Flag, c cli.Command) {
			fmt.Println(usage)
		},
	})

	cmd := c.Main(mainCommand)

	cmd.AddFlag(&cli.Flag{
		Name:     "artifacts",
		Short:    "-a",
		Long:     "--artifacts",
		Argument: true,
		Default:  ".",
	})

	cmd.AddFlag(&cli.Flag{
		Name:     "objects",
		Short:    "-o",
		Long:     "--objects",
		Argument: true,
		Default:  ".",
	})

	cmd.AddFlag(&cli.Flag{
		Name:     "manifest",
		Short:    "-m",
		Long:     "--manifest",
		Argument: true,
		Default:  ".thrall.yml",
	})

	cmd.AddFlag(&cli.Flag{
		Name:     "stage",
		Short:    "-s",
		Long:     "--stage",
		Argument: true,
	})

	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
