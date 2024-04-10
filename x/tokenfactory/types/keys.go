package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "tokenfactory"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokenfactory"
)

var (
	ParamsKey = collections.NewPrefix("p_tokenfactory")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
