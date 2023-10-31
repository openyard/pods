package example

import (
	"github.com/openyard/pods/lamportclock"
	"github.com/openyard/pods/vval"
)

type server struct {
	mvccStore vval.Store
	clock     *lamportclock.LamportClock
}

func newServer(store vval.Store) *server {
	return &server{
		mvccStore: store,
		clock:     lamportclock.New(1),
	}
}
func (s *server) Write(key, value string, requestTS int32) (int32, error) {
	// update own clock to reflect causality
	writeAt := s.clock.Tick(requestTS)
	err := s.mvccStore.Put(vval.NewVersionedKey(key, writeAt), value)
	return writeAt, err
}
