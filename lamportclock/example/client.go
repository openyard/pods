package example

import (
	"fmt"
	"log"

	"github.com/openyard/pods/lamportclock"
	"github.com/openyard/pods/vval"
)

type client struct {
	clock            *lamportclock.LamportClock
	server1, server2 *server
}

func newClient() *client {
	store := vval.NewReplicatedKVStore(vval.NewMVCCStore())
	return &client{
		clock:   lamportclock.New(1),
		server1: newServer(store),
		server2: newServer(store),
	}
}

func (c *client) Write() {
	server1WrittenAt, err := c.server1.Write("name", "Alice", c.clock.GetLatestTime())
	if err != nil {
		log.Panicf("err: %s", err)
	}
	c.clock.UpdateTo(server1WrittenAt)

	server2WrittenAt, err := c.server2.Write("title", "Microservices", c.clock.GetLatestTime())
	if err != nil {
		log.Panicf("err: %s", err)
	}
	c.clock.UpdateTo(server2WrittenAt)

	c.assertTrue(server2WrittenAt > server1WrittenAt)
	fmt.Printf("[DEBUG] client clock at: %d", c.clock.GetLatestTime())
}

func (c *client) assertTrue(b bool) {
	if !b {
		log.Panicf("write failure")
	}
}
