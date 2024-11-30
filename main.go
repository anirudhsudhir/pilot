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
		case "container":
			runContainer()
		}
	} else {
		log.Println("[HOST] Arguments -> ", os.Args)
		log.Println("[HOST] Not enough arguments to start the container")
		log.Println("[HOST] Usage: ./pilot run <command> [args]")
	}
}

func initContainerParent() {
	childProcess := "/proc/self/exe"
	log.Printf("[CONTAINER-PARENT] Starting container parent process with PID %d -> %s", os.Getpid(), childProcess)
	cmd := exec.Command(childProcess, append([]string{"container"}, os.Args[2:]...)...)

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
	log.Printf("[CONTAINER] Starting container with PID %d -> %s", os.Getpid(), os.Args[2])

	errCheck(syscall.Sethostname([]byte("container")))
	errCheck(syscall.Chroot("utils/sample-rootfs"))
	errCheck(os.Chdir("/"))
	errCheck(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("[CONTAINER-PARENT] Failed to start container with process %q -> %v", os.Args[2], err)
	} else {
		errCheck(syscall.Unmount("proc", 0))
	}

	log.Println("[CONTAINER-PARENT] Stopping container")
}

func errCheck(err error) {
	if err != nil {
		log.Fatalln("Encountered error -> ", err)
	}
}
