package executor

import "fmt"

// SSHExecutor is an implementation of the Executor interface that runs
// commands on a remote machine via SSH.
type SSHExecutor struct{}

// NewSSHExecutor creates a new SSHExecutor.
func NewSSHExecutor() *SSHExecutor {
	return &SSHExecutor{}
}

// Execute runs the given command on the remote machine via SSH.
func (e *SSHExecutor) Execute(command, host, user string, key []byte) (string, error) {
	fmt.Printf("Simulating executing command on remote machine via SSH: %s@%s\n", user, host)
	// TODO: Implement code that will execute via SSH
	return "", nil
}
