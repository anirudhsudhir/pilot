package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) >= 3 && os.Args[1] == "run" {
		initContainer()
	} else {
		log.Println("[HOST] Arguments -> ", os.Args)
		log.Println("[HOST] Not enough arguments to start the container")
		log.Println("[HOST] Usage: ./pilot run <command> [args]")
	}
}

func initContainer() {
	log.Println("[CONTAINER-PARENT] Starting container process -> ", os.Args[2])
	cmd := exec.Command(os.Args[2])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("[CONTAINER-PARENT] Failed to start container process %q -> %v", os.Args[2], err)
	}

	log.Println("[CONTAINER-PARENT] Finished running container process")
}
