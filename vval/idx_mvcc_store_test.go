package vval_test

import (
	"testing"

	"github.com/openyard/pods/vval"
)

func TestIndexedMVCCStore_GetRange(t *testing.T) {
	s := vval.NewIdxMVCCStore()
	_ = s.Put("key", "val1")
	_ = s.Put("key", "val2")
	_ = s.Put("key", "val3")
	_ = s.Put("key", "val5")

	assertRange(t, "case-1", []string{"val1", "val2", "val3", "val5"}, s.GetRange("key", 0, 5))
	assertRange(t, "case-2", []string{"val1", "val2", "val3", "val5"}, s.GetRange("key", 1, 5))
	assertRange(t, "case-3", []string{"val1"}, s.GetRange("key", 0, 1))
	assertRange(t, "case-4", []string{}, s.GetRange("key", 5, 6))
	assertRange(t, "case-5", []string{}, s.GetRange("key", 0, 0))
	assertRange(t, "case-6", []string{}, s.GetRange("key", 6, 6))

}

func assertRange(t *testing.T, name string, expected []string, actual []string) {
	if len(expected) != len(actual) {
		t.Errorf("%s: expected (%+v) != actual (%+v)", name, expected, actual)
	}
}
