vhost-index
-----------

A simple statically-linked go application that will display a list
of local containers that have the `VIRTUAL_HOST` environment variable
set, a link to go straight to that vhost, and the status of the
container.

![container list](https://raw.githubusercontent.com/tehbilly/dockerfiles/master/vhost-index/container-list.png)

This was made to scratch my own it, but it's not far off in functionality from being
a basic administration hub. Right now I use [docker-gen](https://github.com/jwilder/docker-gen)
and nginx to route to containers by setting a `VIRTUAL_HOST` environment variable when
running the container. This works and takes a lot of administrative pain off, but
I have something in the works to make it easier/smarter. Until then, I like to see
what I have running and where.

Usage
=====

You will need to mount the docker daemon socket into the container at `/docker.sock`:

```bash
$ docker run -d -p 3000:3000 -v /var/run/docker.sock:/docker.sock tehbilly/vhost-index
```

If you want to use a remote docker daemon, set the `DOCKER_ENDPOINT`
environment variable:

```bash
$ docker run -d -p 3000:3000 -e DOCKER_ENDPOINT=http://localhost:4243 tehbilly/vhost-index
```

By default, the index will generate generate a random password for the user `vadmin`. The
authentication user and password will be in the container's logs. To disable authentication
set the environment variable `AUTH_OFF` to `true`. You can also specify the auth credentials
via `AUTH_USER` and `AUTH_PASS` environment variables.

Future Plans
============

My goal is to have this able to handle 80% of your container information/control needs,
and provide a bundled way to make your containers visible to the outside world. I like
what docker-gen does, but I personally miss the ability to change things at runtime.

Immediate goals:

- Provide pull/tag/push options for images to repositories.
- Show on image/container list which images have new versions available remotely.
- Show docker events as growl-like notifications.
- Stream container logs when viewing container info.
- Tie in to docker's new stats endpoint.
