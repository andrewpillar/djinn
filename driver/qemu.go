package driver

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/runner"
)

type QEMU struct {
	io.Writer

	*SSH

	pidfile string
	process *os.Process

	dir     string
	image   string
	arch    string
	cpus    int
	memory  int
	hostfwd string
}

func (d *QEMU) Create(c context.Context, env []string, objects runner.Passthrough, p runner.Placer) error {
	fmt.Fprintf(d.Writer, "Running with QEMU driver...\n")

	pidfile, err := ioutil.TempFile("", "thrall-qemu-")

	if err != nil {
		return err
	}

	d.pidfile = pidfile.Name()

	bin := fmt.Sprintf("qemu-system-%s", d.arch)
	arg := []string{
		"-daemonize",
		"-display",
		"none",
		"-pidfile",
		d.pidfile,
		"-smp",
		fmt.Sprintf("%d", d.cpus),
		"-m",
		fmt.Sprintf("%d", d.memory),
		"-net",
		"nic,model=virtio",
		"-net",
		"user,hostfwd=tcp:" + d.hostfwd + "-:22",
		"-drive",
		"file=" + filepath.Join(d.dir, d.arch, d.image) + ",media=disk,snapshot=on,if=virtio",
	}

	fmt.Fprintf(d.Writer, "Booting machine with image %s...\n", d.image)

	cmd := exec.Command(bin, arg...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	b, err := ioutil.ReadAll(pidfile)

	if err != nil {
		return err
	}

	trim := strings.Trim(string(b), "\n")
	pid, err := strconv.ParseInt(trim, 10, 64)

	if err != nil {
		return err
	}

	d.process, err = os.FindProcess(int(pid))

	if err != nil {
		return err
	}

	if d.SSH == nil {
		return errors.New("SSH driver for QEMU not initialized")
	}

	if err := d.SSH.Create(c, env, runner.NewPassthrough(), p); err != nil {
		return err
	}

	fmt.Fprintf(d.Writer, "Established SSH connection to machine...\n\n")

	return d.placeObjects(objects, p)
}

func (d *QEMU) Execute(j *runner.Job, c runner.Collector) {
	d.SSH.Execute(j, c)
}

func (d *QEMU) Destroy() {
	d.SSH.Destroy()

	if d.process != nil {
		d.process.Kill()
	}

	if d.pidfile != "" {
		os.Remove(d.pidfile)
	}
}
