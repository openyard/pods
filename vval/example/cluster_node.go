package example

import "github.com/openyard/pods/vval"

type clusterNode struct {
	mvccStore vval.Store
	isLeader  func() bool
	peers     []*clusterNode
}

func newClusterNode() *clusterNode {
	return &clusterNode{
		mvccStore: nil,
		isLeader:  nil,
		peers:     nil,
	}
}
