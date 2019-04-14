package config

import (
	"io"

	"github.com/andrewpillar/thrall/errors"

	"github.com/pelletier/go-toml"
)

type Server struct {
	Assets string

	Net struct {
		Listen string

		SSL struct {
			Listen string
			Cert   string
			Key    string
		}
	}

	Crypto struct {
		Hash string
		Key  string
	}

	Database struct {
		Addr     string
		Name     string
		Username string
		Password string
	}

	Redis struct {
		Addr     string
		Password string
		Idle     int
	}

	Log struct {
		Level  string
		File   string
		Access bool
	}
}

func DecodeServer(r io.Reader) (Server, error) {
	dec := toml.NewDecoder(r)

	server := Server{}
	server.Redis.Idle = 10

	server.Net.Listen = ":80"
	server.Net.SSL.Listen = ":443"

	server.Log.Level = "info"
	server.Log.File = "/dev/stdout"

	if err := dec.Decode(&server); err != nil {
		return server, errors.Err(err)
	}

	return server, nil
}
