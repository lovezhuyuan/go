package main

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {

	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
}
func main() {
	exec.Command("echo", "-n", "my")
}
