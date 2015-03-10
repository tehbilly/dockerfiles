package main

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
	"github.com/k0kubun/pp"
)

var basicTemplate string = `upstream {{.Upstream}} {
	server {{.Ip}}:{{.Port}};
}

server {
	gzip_types text/plain text/css application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript;

	server_name {{.VirtualHost}};
	proxy_buffering off;
	error_log  /proc/self/fd/2;
	access_log /proc/self/fd/1;

	location / {
		proxy_pass http://{{.Upstream}};
		proxy_set_header Host $http_host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme;

		# HTTP 1.1 support
		proxy_http_version 1.1;
		proxy_set_header Connection "";
	}
}
`

func main() {
	client, err := docker.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, c := range containers {
		_, err := client.InspectContainer(c.ID)
		if err != nil {
			fmt.Println("Error inspecting container:", err)
			continue
		}

		pp.Println(c)
	}
}
