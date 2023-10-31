package vval

import "log"

// ReplicatedKVStore is a replicated key-value store
type ReplicatedKVStore struct {
	version       int32
	MVCCStore     *MVCCStore
	replicatedLog ReplicatedLogFunc
}

// NewReplicatedKVStore initializes a new replicated key-value store with given multi-version concurrency control store
func NewReplicatedKVStore(mvccStore *MVCCStore) *ReplicatedKVStore {
	return &ReplicatedKVStore{
		version:   0,
		MVCCStore: mvccStore,
		replicatedLog: func(cmd any) error {
			return nil
		},
	}
}

// WithReplicatedLogFunc return a replicated key-value store with given func for a replicated log
func (rs *ReplicatedKVStore) WithReplicatedLogFunc(f ReplicatedLogFunc) *ReplicatedKVStore {
	rs.replicatedLog = f
	return rs
}

// Put proposes the given key and value to the replicated log and returns an error in case of failure
func (rs *ReplicatedKVStore) Put(key any, value string) error {
	err := make(chan error, 1)
	go func() {
		err <- rs.replicatedLog.Propose(&setValueCommand{key.(*VersionedKey).key, value})
	}()
	return <-err
}

func (rs *ReplicatedKVStore) applySetValueCommand(cmd setValueCommand) (int32, error) {
	log.Printf("[INFO] %T setting key value %+v", rs, cmd)
	rs.version++
	err := rs.MVCCStore.Put(&VersionedKey{cmd.GetKey(), rs.version}, cmd.GetValue())
	return rs.version, err
}

type setValueCommand struct {
	key, value string
}

func (c *setValueCommand) GetKey() string {
	return c.key
}
func (c *setValueCommand) GetValue() string {
	return c.value
}
