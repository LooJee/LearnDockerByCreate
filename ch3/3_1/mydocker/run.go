package main

import (
	"LearnDockerByCreate/ch3/3_1/mydocker/container"
	log "github.com/sirupsen/logrus"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
