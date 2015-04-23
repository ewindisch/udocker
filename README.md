udocker
-------

Run docker programs as untrusted local users.

Users cannot configure volume mounts, enable privileged operations,
or do other "unsafe" things with Docker. They cannot explicitly pull
images, save images, or do anything else that may be unsafe.

Note that this program is not **completely** safe for running Docker as an
unprivileged user. It is not possible to prevent the user from
implicitly causing a 'docker pull' of malicious images, which may use
setuid binaries to provide local privilege escalation (note: this
would provide escalation to the "limited" root user typical of a
Docker container); Remediation of this would include preventing
access to Docker registries and eventually incorporating a mechanism
to disallow 'docker run' from implicitly pulling an image.

Usage
-----

On the same host as a running Docker daemon:

$ udocker image [args...]

Installation
------------

This depends the binary being installed as setgid 'docker',
allowing access to the Docker socket.

$ go build udocker.go
$ chgrp docker udocker
$ chmod g+s udocker
$ mv udocker /usr/local/bin

