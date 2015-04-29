udocker
-------

Run docker containers as untrusted local users, allowing users to
run containers (only) under their current UID.

Users cannot configure volume mounts, enable privileged operations,
or do other "unsafe" things with Docker. They cannot explicitly pull
images, save images, or do anything else that may be unsafe. For
safety purposes, all privileged capabilities are disabled, including
NET_RAW and SETUID. Users should not expect to be able to use
ICMP ping or setuid binaries from their images.

Usage
-----

On the same host as a running Docker daemon:

```
$ udocker image [args...]
```

Installation
------------

This depends the binary being installed as setgid 'docker',
allowing access to the Docker socket.

```
$ go build udocker.go
$ chgrp docker udocker
$ chmod g+s udocker
$ mv udocker /usr/local/bin
```

Example
--------

```
$ id
uid=1000(eric) gid=1000(eric)
groups=1000(eric)

$ udocker busybox id
iduid=1000(default) gid=1000(default)

$ udocker busybox echo "hello world"
hello world
```
