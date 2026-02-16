package system

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func FlushDNSCache() error {
	if runtime.GOOS != "darwin" {
		return nil
	}

	if err := RunCommand("sudo", "dscacheutil", "-flushcache"); err != nil {
		return err
	}

	_ = RunCommand("sudo", "killall", "-HUP", "mDNSResponder")
	return nil
}

func CloseParentTerminal() {
	ppid := os.Getppid()
	if ppid <= 1 {
		return
	}

	proc, err := os.FindProcess(ppid)
	if err != nil {
		return
	}

	if err := proc.Signal(syscall.SIGHUP); err != nil && !errors.Is(err, os.ErrProcessDone) {
		_ = proc.Signal(syscall.SIGTERM)
	}
}
