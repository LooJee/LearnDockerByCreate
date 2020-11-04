package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

	/*
		MS_NOEXEC: 在本文件系统中不允许运行其它程序
		MS_NOSUID: 在本系统中运行程序时，不允许 set-user-ID 或 set-group-ID
	*/
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	//在容器内部挂载 proc ，后面查看进程信息时，看到的是容器的进程信息，而不是外部的容器信息
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}
