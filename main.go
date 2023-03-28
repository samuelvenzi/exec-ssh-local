package main

import (
	"exec-ssh-local/executor"
)

func main() {
	test := true

	var goexec executor.Executor
	if test {
		// Instantiate local executor
		goexec = executor.NewLocalExecutor()
	} else {
		// Instantiate SSH executor
		goexec = executor.NewSSHExecutor("host", "ubuntu", []byte("key"))
	}
	_, _ = goexec.Execute("command")
}
