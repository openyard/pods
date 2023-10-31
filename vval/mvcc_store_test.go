package vval_test

import (
	"testing"

	"github.com/openyard/pods/vval"
)

func TestMVCCStore_Get(t *testing.T) {
	s := vval.NewMVCCStore()
	_ = s.Put(vval.NewVersionedKey("key", 1), "val1")
	_ = s.Put(vval.NewVersionedKey("key", 2), "val2")
	_ = s.Put(vval.NewVersionedKey("key", 3), "val3")
	_ = s.Put(vval.NewVersionedKey("key", 5), "val5")

	assertEquals(t, "case-1", "val2", s.Get("key", 2))
	assertEquals(t, "case-2", "val5", s.Get("key", 7))
	assertEquals(t, "case-3", "val3", s.Get("key", 4))
	assertEquals(t, "case-4", "", s.Get("other-key", 1))
}

func assertEquals(t *testing.T, name string, expected string, actual string) {
	if expected != actual {
		t.Errorf("%s: expected (%s) != actual (%s)", name, expected, actual)
	}
}
