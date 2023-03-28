package executor

import "fmt"

// SSHExecutor is an implementation of the Executor interface that runs
// commands on a remote machine via SSH.
type SSHExecutor struct {
	host string
	user string
	key  []byte
}

// NewSSHExecutor creates a new SSHExecutor.
func NewSSHExecutor(host, user string, key []byte) *SSHExecutor {
	return &SSHExecutor{
		host: host,
		user: user,
		key:  key,
	}
}

// Execute runs the given command on the remote machine via SSH.
func (e *SSHExecutor) Execute(command string) (string, error) {
	fmt.Printf("Simulating executing command on remote machine via SSH: %s@%s\n", e.user, e.host)
	// TODO: Implement code that will execute via SSH
	return "", nil
}
