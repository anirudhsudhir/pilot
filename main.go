package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) >= 3 {
		switch os.Args[1] {
		case "run":
			initContainerParent()
			break
		case "container":
			runContainer()
			break
		}
	} else {
		log.Println("[HOST] Arguments -> ", os.Args)
		log.Println("[HOST] Not enough arguments to start the container")
		log.Println("[HOST] Usage: ./pilot run <command> [args]")
	}
}

func initContainerParent() {
	log.Println("[CONTAINER-PARENT] Starting container parent process -> ", "/proc/self/exe")
	cmd := exec.Command("/proc/self/exe", append([]string{"container"}, os.Args[2:]...)...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("[CONTAINER-PARENT] Failed to start container parent process %q -> %v", "/proc/self/exe", err)
	}

	log.Println("[CONTAINER-PARENT] Stopping container parent process")
}

func runContainer() {
	log.Println("[CONTAINER] Starting container with process -> ", os.Args[2])

	syscall.Sethostname([]byte("container"))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("[CONTAINER-PARENT] Failed to start container with process %q -> %v", os.Args[2], err)
	}

	log.Println("[CONTAINER-PARENT] Stopping container")
}
