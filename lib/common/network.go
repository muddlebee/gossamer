// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package common

import (
	ma "github.com/multiformats/go-multiaddr"
)

// Health is network information about host needed for the rpc server
type Health struct {
	Peers           int
	IsSyncing       bool
	ShouldHavePeers bool
}

// NetworkState is network information about host needed for the rpc server and the runtime
type NetworkState struct {
	PeerID     string
	Multiaddrs []ma.Multiaddr
}

// PeerInfo is network information about peers needed for the rpc server
type PeerInfo struct {
	PeerID     string
	Roles      Roles
	BestHash   Hash
	BestNumber uint64
}

// Roles is the type of node.
type Roles byte

const (
	// NoNetworkRole runs a node without networking
	NoNetworkRole Roles = 0
	// FullNodeRole allow you to read the current state of the chain and to submit and validate
	// extrinsics directly on the network without relying on a centralised infrastructure provider.
	FullNodeRole Roles = 1
	// LightClientRole node has only the runtime and the current state, but does not store past
	// blocks and so cannot read historical data without requesting it from a node that has it.
	LightClientRole Roles = 2
	// AuthorityRole runs the node as a block-producing and finalising node
	AuthorityRole Roles = 4
)
