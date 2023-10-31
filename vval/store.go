package vval

// Store interface...
type Store interface {
	Put(key any, value string) error
}
