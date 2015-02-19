Hekad Log Shipper
-----------------

A simple container that will ship logs from docker containers to elasticsearch (or anywhere,
really, considering the hekad instance is configurable).

```bash
$ docker run -d --name heka-dockerlogs --link elasticsearch:es -v /var/run/docker.sock:/var/run/docker.sock tehbilly/heka-dockerlogs
```

If you want to configure this beyond shipping to a linked elasticsearch instance, mount your
configuration directory do `/conf` in the container. The rest you can find in heka's docs.
