vhost-index
-----------

A simple statically-linked go application that will display a list
of local containers that have the `VIRTUAL_HOST` environment variable
set, a link to go straight to that vhost, and the status of the
container.

Why? Why not! The image clocks in at about 6.2MB, and it's handy for me.
The use case for me is listing all containers I publish with
[docker-gen](https://github.com/jwilder/docker-gen) and nginx.

Usage
=====

You need to expose the port the container listens on:

```
$ docker run -d -p 8080:80 -v /var/run/docker.sock:/docker.sock tehbilly/vhost-index
```

If you want to use a remote docker daemon, set the `DOCKER_ENDPOINT`
environment variable:

```
$ docker run -d -p 8080:80 -e DOCKER_ENDPOINT=http://localhost:4243 tehbilly/vhost-index
```

Ugly
====

Yep, it's ugly. I threw this together quickly out of boredom. I plan to make
it pretty and provide more useful information. The only other planned features are:

- Establish basic HTTP auth protection with `HTTP_USER` and `HTTP_PASS` variables.
- Enable stopping/starting containers.
