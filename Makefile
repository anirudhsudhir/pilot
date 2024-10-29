all: pilot
	sudo ./pilot run /bin/bash

pilot: main.go
	go build
