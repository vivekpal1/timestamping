package types

const (
	// ModuleName defines the module name
	ModuleName = "btctimestamping"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_btctimestamping"
)

var (
	ParamsKey = []byte("p_btctimestamping")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
