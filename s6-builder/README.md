# tehbilly/s6-builder

This image is used to build [s6](http://www.skarnet.org/software/s6/) linked statically against
[musl](http://www.musl-libc.org/). The binaries (and libs) are built from an Arch OS to utilize
the wonderful `AUR` packages, greatly simplifying the entire process.

Images will be tagged with the `s6` release version number, with `latest` on the newest version.

## Usage

Clone the repository, and get 'er done.

```bash
$ git clone https://github.com/tehbilly/dockerfiles tb-dockerfiles
$ cd tb-dockerfiles/s6-builder
$ docker build -t $USER/s6-builder .
$ mkdir output
$ docker run --rm -v $(pwd)/output:/output $USER/s6-builder
```

Afterwards you'll have a tarball in the mounted volume.
