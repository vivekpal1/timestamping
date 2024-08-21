package types

const (
	// ModuleName defines the module name
	ModuleName = "timestamping"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_timestamping"
)

var (
	ParamsKey = []byte("p_timestamping")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
