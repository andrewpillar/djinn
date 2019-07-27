package driver

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/runner"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Docker struct {
	io.Writer

	client *client.Client
	volume types.Volume

	env        []string
	containers []string

	mutex *sync.Mutex

	image      string
	workspace  string
}

func NewDocker(w io.Writer, image, workspace string) *Docker {
	return &Docker{
		Writer:     w,
		mutex:      &sync.Mutex{},
		image:      image,
		workspace:  workspace,
	}
}

func (d *Docker) Create(c context.Context, env []string, objects runner.Passthrough, p runner.Placer) error {
	fmt.Fprintf(d.Writer, "Running with Docker driver...\n")

	var err error

	d.client, err = client.NewEnvClient()

	if err != nil {
		return err
	}

	done := make(chan struct{})
	errs := make(chan error)

	go func() {
		d.volume, err = d.client.VolumeCreate(c, volume.VolumeCreateBody{})

		if err != nil {
			errs <- err
			return
		}

		rc, err := d.client.ImagePull(c, d.image, types.ImagePullOptions{})

		if err != nil {
			errs <- err
			return
		}

		defer rc.Close()

		io.Copy(ioutil.Discard, rc)

		done <- struct{}{}
	}()

	select {
	case <-c.Done():
		return c.Err()
	case <-done:
		break
	case err = <-errs:
		return err
	}

	image, _, err := d.client.ImageInspectWithRaw(c, d.image)

	if err != nil {
		return err
	}

	fmt.Fprintf(d.Writer, "Using Docker image %s - %s...\n\n", d.image, image.ID)

	d.env = env

	return d.placeObjects(objects, p)
}

func (d *Docker) Execute(j *runner.Job, c runner.Collector) {
	hostCfg := &container.HostConfig{
		Mounts: []mount.Mount{
			mount.Mount{
				Type:   mount.TypeVolume,
				Source: d.volume.Name,
				Target: d.workspace,
			},
		},
	}

	cfg := &container.Config{
		Image: d.image,
		Cmd:   []string{"true"},
	}

	ctx := context.Background()

	ctr, err := d.client.ContainerCreate(ctx, cfg, hostCfg, nil, "")

	if err != nil {
		j.Failed(err)
		return
	}

	d.mutex.Lock()
	d.containers = append(d.containers, ctr.ID)
	d.mutex.Unlock()

	script := j.Name + ".sh"
	buf := createScript(j)

	header := &tar.Header{
		Typeflag: tar.TypeReg,
		Name:     "/bin/" + script,
		Size:     int64(buf.Len()),
		Mode:     755,
	}

	pr, pw := io.Pipe()
	defer pr.Close()

	tw := tar.NewWriter(pw)

	go func() {
		defer tw.Close()
		defer pw.Close()

		tw.WriteHeader(header)
		io.Copy(tw, buf)
	}()

	d.client.CopyToContainer(ctx, ctr.ID, d.workspace, pr, types.CopyToContainerOptions{})

	cfg.Cmd = []string{}
	cfg.Env = d.env
	cfg.Entrypoint = []string{script}

	ctr, err = d.client.ContainerCreate(ctx, cfg, hostCfg, nil, "")

	if err != nil {
		j.Failed(err)
		return
	}

	d.mutex.Lock()
	d.containers = append(d.containers, ctr.ID)
	d.mutex.Unlock()

	if err := d.client.ContainerStart(ctx, ctr.ID, types.ContainerStartOptions{}); err != nil {
		j.Failed(err)
		return
	}

	logOpts := types.ContainerLogsOptions{
		ShowStderr: true,
		ShowStdout: true,
		Timestamps: false,
		Follow:     true,
	}

	go func() {
		rc, err := d.client.ContainerLogs(ctx, ctr.ID, logOpts)

		if err != nil {
			if err == io.EOF {
				return
			}

			return
		}

		defer rc.Close()

		stdcopy.StdCopy(j.Writer, ioutil.Discard, rc)
	}()

	status, errs := d.client.ContainerWait(ctx, ctr.ID, container.WaitConditionNotRunning)
	code := 0

	select {
		case err := <-errs:
			if err != nil {
				j.Failed(err)
				return
			}
		case resp := <-status:
			code = int(resp.StatusCode)
	}

	if code != 0 {
		j.Failed(nil)
	} else {
		j.Status = runner.Passed
	}

	if len(j.Artifacts) > 0 {
		fmt.Fprintf(j.Writer, "\n")
	}

	for src, dst := range j.Artifacts {
		fmt.Fprintf(j.Writer, "Collecting artifact %s => %s\n", src, dst)

		rc, _, err := d.client.CopyFromContainer(ctx, ctr.ID, src)

		if err != nil {
			fmt.Fprintf(
				j.Writer,
				"Failed to collect artifact %s => %s: %s\n",
				src,
				dst,
				errors.Cause(err),
			)
			continue
		}

		defer rc.Close()

		tr := tar.NewReader(rc)

		for {
			header, err := tr.Next()

			if err != nil {
				if err == io.EOF {
					break
				}

				fmt.Fprintf(j.Writer, "Failed to collect artifact %s => %s: %s\n", src, dst, err)
				break
			}

			switch header.Typeflag {
				case tar.TypeDir:
					break
				case tar.TypeReg:
					if _, err := c.Collect(dst, tr); err != nil {
						fmt.Fprintf(
							j.Writer,
							"Failed to collect artifact %s => %s: %s\n",
							src,
							dst,
							errors.Cause(err),
						)
					}
			}
		}
	}

	if j.Status == runner.Failed {
		j.Failed(nil)
	}
}

func (d *Docker) Destroy() {
	if d.client == nil {
		return
	}

	ctx := context.Background()

	opts := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	for _, ctr := range d.containers {
		d.client.ContainerRemove(ctx, ctr, opts)
	}

	d.client.VolumeRemove(ctx, d.volume.Name, true)
}

func (d *Docker) placeObjects(objects runner.Passthrough, p runner.Placer) error {
	if len(objects) == 0 {
		return nil
	}

	cfg := &container.Config{
		Image: d.image,
		Cmd:   []string{"true"},
	}

	hostCfg := &container.HostConfig{
		Mounts: []mount.Mount{
			mount.Mount{
				Type:   mount.TypeVolume,
				Source: d.volume.Name,
				Target: d.workspace,
			},
		},
	}

	ctx := context.Background()

	ctr, err := d.client.ContainerCreate(ctx, cfg, hostCfg, nil, "")

	if err != nil {
		return err
	}

	for src, dst := range objects {
		fmt.Fprintf(d.Writer, "Placing object %s => %s\n", src, dst)

		info, err := p.Stat(src)

		if err != nil {
			fmt.Fprintf(d.Writer, "Failed to place object %s => %s: %s\n", src, dst, errors.Cause(err))
			continue
		}

		header, err := tar.FileInfoHeader(info, info.Name())

		if err != nil {
			fmt.Fprintf(d.Writer, "Failed to place object %s => %s: %s\n", src, dst, errors.Cause(err))
			continue
		}

		header.Name = strings.TrimPrefix(dst, d.workspace)

		pr, pw := io.Pipe()
		defer pr.Close()

		tw := tar.NewWriter(pw)

		go func(src string) {
			defer tw.Close()
			defer pw.Close()

			tw.WriteHeader(header)

			if _, err := p.Place(src, tw); err != nil {
				fmt.Fprintf(d.Writer, "Failed to place object %s => %s: %s\n", src, dst, errors.Cause(err))
			}
		}(src)

		err = d.client.CopyToContainer(ctx, ctr.ID, d.workspace, pr, types.CopyToContainerOptions{})

		if err != nil {
			fmt.Fprintf(d.Writer, "Failed to place object %s => %s: %s\n", src, dst, errors.Cause(err))
		}
	}

	d.client.ContainerRemove(ctx, ctr.ID, types.ContainerRemoveOptions{})

	fmt.Fprintf(d.Writer, "\n")

	return nil
}
