package main

import (
	"os"
	"os/exec"
	"strconv"
)

func main() {
	uid := strconv.Itoa(os.Getuid())

	cmdargs := append([]string{"run", "-u", uid, "--rm", "-i", "--"}, os.Args[1:]...)
	cmd := exec.Command("docker", cmdargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
