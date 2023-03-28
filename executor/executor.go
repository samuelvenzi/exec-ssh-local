package executor

// Executor is an interface that defines the methods that an executor must
// implement.
type Executor interface {
	// Execute runs the given command on the machine.
	Execute(command string) (string, error)
}
