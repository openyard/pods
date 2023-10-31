package vval

// ReplicatedLogFunc ...
type ReplicatedLogFunc func(cmd any) error

// Propose given command
func (rlf ReplicatedLogFunc) Propose(cmd any) error {
	return rlf(cmd)
}
