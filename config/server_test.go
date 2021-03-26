package config

import (
	"strings"
	"testing"
)


func Test_DecodeServer(t *testing.T) {
	r := strings.NewReader(`

host "https://djinn-di.com"

log debug "/dev/stdout"

drivers [
	"docker",
	"qemu-x86_64",
]

net {
	listen "localhost:8080"

	ssl {
		cert "/var/lib/ssl/server.crt"
		key  "/var/lib/ssl/server.key"
	}
}

crypto {
	hash  "123456"
	block "123456"
	salt  "123456"
	auth  "123456"
}

smtp {
	addr "smtp.example.com:587"

	ca "/etc/ssl/cert.pem"

	admin "noreply@djinn-ci.com"

	username "postmaster@example.com"
	password "secret"
}

database {
	addr "localhost:5432"
	name "djinn"

	username "djinn"
	password "secret"
}

redis {
	addr "localhost:6379"
}

store images {
	type "file"
	path "/var/lib/djinn/images"
}

store artifacts {
	type "file"
	path "/var/lib/djinn/artifacts"
}

store objects {
	type "file"
	path  "/var/lib/djinn/objects"
	limit 5242880
}

provider github {
	secret        "123456"
	client_id     "123456" 
	client_secret "123456"
}

provider gitlab {
	secret        "123456"
	client_id     "123456"
	client_secret "123456"
}`)

	p := newParser(t.Name(), r, func(name string, line, col int, msg string) {
		t.Errorf("%s,%d:%d - %s\n", name, line, col, msg)
	})
	p.parse()

	if err := p.err(); err != nil {
		t.Fatal(err)
	}
}
