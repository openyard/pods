package vval

import (
	"log"
	"slices"
	"sync"
)

type IdxMVCCStore struct {
	sync.RWMutex
	version int32
	kvIdx   map[string][]int32
	kv      map[*VersionedKey]string
}

// NewIdxMVCCStore initializes a new indexed multi-version concurrency control store
func NewIdxMVCCStore() *IdxMVCCStore {
	return &IdxMVCCStore{
		version: 0,
		kvIdx:   make(map[string][]int32),
		kv:      make(map[*VersionedKey]string),
	}
}

func (s *IdxMVCCStore) Put(key, value string) int32 {
	s.Lock()
	defer s.Unlock()

	s.version++
	s.kv[NewVersionedKey(key, s.version)] = value
	s.updateVersionIndex(key, s.version)
	return s.version
}

func (s *IdxMVCCStore) GetRange(key string, fromRev, toRev int32) []string {
	s.Lock()
	defer s.Unlock()

	versions := s.kvIdx[key]
	maxRevForKey := slices.Max(versions)
	var revToRead int32
	if maxRevForKey > toRev {
		revToRead = toRev
	} else {
		revToRead = maxRevForKey
	}

	versionMap, values := subMap(s.kv, NewVersionedKey(key, fromRev), NewVersionedKey(key, revToRead))
	log.Printf("[INFO] %T: available version keys %+v.Reading@.%d:%d", s, versionMap, fromRev, toRev)
	return values
}

func (s *IdxMVCCStore) updateVersionIndex(key string, newVersion int32) {
	versions := s.getVersions(key)
	versions = append(versions, newVersion)
	s.kvIdx[key] = versions
}

func (s *IdxMVCCStore) getVersions(key string) []int32 {
	versions, ok := s.kvIdx[key]
	if !ok {
		versions = make([]int32, 0)
		s.kvIdx[key] = versions
	}
	return versions
}

func subMap(kv map[*VersionedKey]string, from, to *VersionedKey) (map[*VersionedKey]string, []string) {
	subset := make(map[*VersionedKey]string)
	values := make([]string, 0)
	for k, v := range kv {
		if k.CompareTo(from) >= 0 && k.CompareTo(to) <= 0 {
			subset[k] = v
			values = append(values, v)
		}
	}
	return subset, values
}
