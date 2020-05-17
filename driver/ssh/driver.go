// Package ssh provides an implemention of an SSH driver for executing jobs on.
package ssh

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/andrewpillar/thrall/driver"
	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/runner"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

type SSH struct {
	io.Writer

	client  *ssh.Client
	env     []string

	// Addr is the address of the machine to connect to.
	Addr string

	// User is the user to connect via SSH as.
	User string

	// Key is the path to the SSH key to use during SSH.
	Key string

	// Timeout is the timeout to use when attempting the SSH connection.
	Timeout time.Duration
}

var _ runner.Driver = (*SSH)(nil)

// Init initializes a new SSH driver using the given io.Writer, and
// configuration map. Detailed below are the values, types, and default values
// that are used in the configuration map.
//
// Key - The key for the SSH driver is specified via the "key" field. It is
// expected for this to be a string, the default value is $HOME/.ssh/.id_rsa.
//
// Timeout - The timeout for the SSH driver is specified via the "timeout"
// field. It is expected for this to be an int64 for the timeout in seconds. The
// default value is 60.
//
// Address - The address for the SSH driver is specified via the "address"
// field. It is expected for this to be a string, there is no default value
// for this.
func Init(w io.Writer, cfg map[string]interface{}) runner.Driver {
	key, ok := cfg["key"].(string)

	if !ok {
		key = filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
	}

	timeout, ok := cfg["timeout"].(int64)

	if !ok {
		timeout = 60
	}

	addr, _ := cfg["address"].(string)

	return &SSH{
		Writer:  w,
		Addr:    addr,
		User:    "root",
		Key:     key,
		Timeout: time.Duration(time.Second*time.Duration(timeout)),
	}
}

// Create opens up the SSH connection to the remote machine as configured via a
// previous call to Init. The given env slice is used to set an unexported
// variable for setting environment variables during job execution.
func (s *SSH) Create(c context.Context, env []string, objs runner.Passthrough, p runner.Placer) error {
	if s.Writer == nil {
		return errors.New("cannot create driver with nil io.Writer")
	}

	fmt.Fprintf(s.Writer, "Running with SSH driver...\n")

	ticker := time.NewTicker(time.Second)
	after := time.After(s.Timeout)

	client := make(chan *ssh.Client)

	b, err := ioutil.ReadFile(s.Key)

	if err != nil {
		return err
	}

	signer, err := ssh.ParsePrivateKey(b)

	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				cfg := &ssh.ClientConfig{
					User: s.User,
					Auth: []ssh.AuthMethod{
						ssh.PublicKeys(signer),
					},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
					Timeout:         time.Second,
				}

				fmt.Fprintf(s.Writer, "Connecting to %s...\n", s.Addr)

				cli, err := ssh.Dial("tcp", s.Addr, cfg)

				if err != nil {
					break
				}
				client <- cli
			}
		}
	}()

	select {
	case <-c.Done():
		return c.Err()
	case <-after:
		return fmt.Errorf("Timed out trying to connect to %s...\n", s.Addr)
	case cli := <-client:
		s.client = cli
	}

	fmt.Fprintf(s.Writer, "Established SSH connection to %s...\n\n", s.Addr)

	s.env = env
	return s.PlaceObjects(objs, p)
}

// Execute will perform the given runner.Job. This turns it into a shell script
// that is executed on the remote machine once placed via SFTP. Before the job
// is executed however, the environment variables given via Create are set for
// the SSH session being used to invoke the script. The stderr, and stdout
// streams are forwarded to the underlying io.Writer.
func (s *SSH) Execute(j *runner.Job, c runner.Collector) {
	sess, err := s.client.NewSession()

	if err != nil {
		j.Failed(err)
		return
	}

	defer sess.Close()

	script := strings.Replace(j.Name + ".sh", " ", "-", -1)
	buf := driver.CreateScript(j)

	cli, err := sftp.NewClient(s.client)

	if err != nil {
		j.Failed(err)
		return
	}

	defer cli.Close()

	f, err := cli.Create(script)

	if err != nil {
		j.Failed(err)
		return
	}

	io.Copy(f, buf)

	if err := f.Chmod(0755); err != nil {
		j.Failed(err)
		return
	}

	f.Close()

	for _, e := range s.env {
		parts := strings.SplitN(e, "=", 2)

		if len(parts) > 1 {
			if err := sess.Setenv(parts[0], parts[1]); err != nil {
				fmt.Fprintf(j.Writer, "Failed to setenv %s: %s\n", e, err)
			}
		}
	}

	sess.Stdout = j.Writer
	sess.Stderr = j.Writer

	if err := sess.Run("./" + script); err != nil {
		if _, ok := err.(*ssh.ExitError); ok {
			err = nil
		}

		j.Failed(err)
	} else {
		j.Status = runner.Passed
	}

	s.collectArtifacts(j.Writer, j, c)
	cli.Remove(script)
}

// Destroy closes the SSH connection.
func (s *SSH) Destroy() {
	if s.client != nil {
		s.client.Close()
	}
}

func (s *SSH) collectArtifacts(w io.Writer, j *runner.Job, c runner.Collector) {
	if len(j.Artifacts.Values) == 0 {
		return
	}

	cli, err := sftp.NewClient(s.client)

	if err != nil {
		j.Failed(err)
		return
	}

	defer cli.Close()

	fmt.Fprintf(w, "\n")

	for src, dst := range j.Artifacts.Values {
		fmt.Fprintf(w, "Collecting artifact %s => %s\n", src, dst)

		f, err := cli.Open(src)

		if err != nil {
			fmt.Fprintf(
				w,
				"Failed to collect artifact %s => %s: %s\n",
				src,
				dst,
				errors.Cause(err),
			)
			continue
		}

		defer f.Close()

		if _, err := c.Collect(dst, f); err != nil {
			fmt.Fprintf(
				w,
				"Failed to collect artifact %s => %s: %s\n",
				src,
				dst,
				errors.Cause(err),
			)
		}
	}
}

// PlaceObjects copies the given objects from the given placer onto the
// environment via SFTP.
func (s *SSH) PlaceObjects(objects runner.Passthrough, p runner.Placer) error {
	if len(objects.Values) == 0 {
		return nil
	}

	cli, err := sftp.NewClient(s.client)

	if err != nil {
		return err
	}

	defer cli.Close()

	for src, dst := range objects.Values {
		fmt.Fprintf(s.Writer, "Placing object %s => %s\n", src, dst)

		f, err := cli.OpenFile(dst, os.O_WRONLY|os.O_APPEND|os.O_CREATE)

		if err != nil {
			fmt.Fprintf(
				s.Writer,
				"Failed to open object file on guest %s => %s: %s\n",
				src,
				dst,
				errors.Cause(err),
			)
			continue
		}

		defer f.Close()

		if err := f.Chmod(0700); err != nil {
			fmt.Fprintf(
				s.Writer,
				"Failed to chmod 0700 object %s => %s: %s\n",
				src,
				dst,
				errors.Cause(err),
			)
			continue
		}

		if _, err := p.Place(src, f); err != nil {
			fmt.Fprintf(
				s.Writer,
				"Failed to place object %s => %s: %s\n",
				src,
				dst,
				errors.Cause(err),
			)
			continue
		}
	}
	fmt.Fprintf(s.Writer, "\n")
	return nil
}
