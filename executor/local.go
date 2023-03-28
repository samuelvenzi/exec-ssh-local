package executor

import "fmt"

// LocalExecutor is an implementation of the Executor interface that runs
// commands on the local machine.
type LocalExecutor struct{}

// NewLocalExecutor creates a new LocalExecutor.
func NewLocalExecutor() *LocalExecutor {
	return &LocalExecutor{}
}

// Execute runs the given command on the local machine.
func (e *LocalExecutor) Execute(command string) (string, error) {
	fmt.Printf("Simulating executing command on local machine\n")
	// TODO: Implement code that will execute locally
	return "", nil
}
