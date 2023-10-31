package vval

import (
	"sort"
	"sync"
)

// MVCCStore is a multi-version concurrency control store
type MVCCStore struct {
	sync.RWMutex
	entries map[*VersionedKey]string
}

// NewMVCCStore initializes a new multi-version concurrency control store
func NewMVCCStore() *MVCCStore {
	return &MVCCStore{
		entries: make(map[*VersionedKey]string),
	}
}

// Put inserts or updates the given versioned key
func (s *MVCCStore) Put(key *VersionedKey, value string) error {
	s.sync(func() {
		s.entries[key] = value
	})
	return nil
}

// Get returns the value for an entry with the greatest key less than or equal to the given key
func (s *MVCCStore) Get(key string, readAt int32) string {
	entry := s.floorEntry(&VersionedKey{key, readAt})
	if entry != nil {
		return entry.v
	}
	return ""
}

func (s *MVCCStore) floorEntry(key *VersionedKey) *entry {
	var nearEntry *VersionedKey
	s.sync(func() {
		keys := make(vkSlice, 0, len(s.entries))
		for k := range s.entries {
			keys = append(keys, k)
		}
		sort.Sort(keys)

		for _, k := range keys {
			if key.GetKey() != k.GetKey() {
				continue
			}
			if key.CompareTo(k) == 1 || key.CompareTo(k) == 0 {
				nearEntry = k
			}
		}
	})
	if nearEntry == nil {
		return nil
	}
	return &entry{nearEntry, s.entries[nearEntry]}
}

func (s *MVCCStore) sync(f func()) {
	s.Lock()
	defer s.Unlock()
	f()
}

type entry struct {
	k *VersionedKey
	v string
}

type vkSlice []*VersionedKey

// Len is part of sort.Interface.
func (d vkSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d vkSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d vkSlice) Less(i, j int) bool {
	return d[i].CompareTo(d[j]) == -1
}
