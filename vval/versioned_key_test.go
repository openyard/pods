package vval_test

import (
	"testing"

	"github.com/openyard/pods/vval"
)

func TestVersionedKey_Compare(t *testing.T) {
	k1 := vval.NewVersionedKey("alice", 1)
	k2 := vval.NewVersionedKey("alice", 1)
	k3 := vval.NewVersionedKey("alice", 2)
	k4 := vval.NewVersionedKey("bob", 1)

	assert(t, "case-1", 0, k1.CompareTo(k2))
	assert(t, "case-2", -1, k1.CompareTo(k3))
	assert(t, "case-3", -1, k1.CompareTo(k4))
	assert(t, "case-4", -1, k2.CompareTo(k3))
	assert(t, "case-4", 1, k3.CompareTo(k2))
	assert(t, "case-5", 1, k4.CompareTo(k1))
}

func assert(t *testing.T, name string, expected, actual int) {
	if expected != actual {
		t.Errorf("%s: expected (%d) != actual (%d)", name, expected, actual)
	}
}
