package vval

// VersionedKey provides a key with a version and implements a Comparer interface to allow natural ordering
type VersionedKey struct {
	key     string
	version int32
}

// NewVersionedKey returns a initialized versioned key with given key and version
func NewVersionedKey(key string, version int32) *VersionedKey {
	return &VersionedKey{key, version}
}

// GetKey returns the key of the versioned key
func (vk *VersionedKey) GetKey() string {
	return vk.key
}

// GetVersion returns the version of the versioned key
func (vk *VersionedKey) GetVersion() int32 {
	return vk.version
}

// CompareTo compares two versioned keys and returns
// -1 if vk LT o
// 0 if vk EQ o
// 1 if vk GT o
func (vk *VersionedKey) CompareTo(o *VersionedKey) int {
	b := vk.key == o.key
	if !b {
		if vk.key > o.key {
			return 1
		}
		if vk.key < o.key {
			return -1
		}
	}

	if vk.version > o.version {
		return 1
	}
	if vk.version < o.version {
		return -1
	}
	return 0
}

// Comparer interface provides a CompareTo func
type Comparer interface {
	CompareTo(b *VersionedKey) int
}
