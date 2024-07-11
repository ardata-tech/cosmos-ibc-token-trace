package core

var chainEndpoints = map[string]string{
	"neutron-1":   "https://rest.cosmos.directory/neutron/cosmos",
	"osmosis-1":   "https://rest.cosmos.directory/osmosis/cosmos",
	"phoenix-1":   "https://rest.cosmos.directory/terra2/cosmos",
	"stargaze-1":  "https://rest.cosmos.directory/stargaze/cosmos",
	"cosmoshub-4": "https://rest.cosmos.directory/cosmoshub/cosmos",
	"stride-1":    "https://rest.cosmos.directory/stride/cosmos",
}

var originalDenoms = map[string]string{
	"dAsset": "factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/dAsset",
	"lAsset": "factory/neutron1lzecpea0qxw5xae92xkm3vaddeszr278k7w20c/lAsset",
}

func GetChainEndpoints() map[string]string {
	return chainEndpoints
}

func GetOriginalDenoms() map[string]string {
	return originalDenoms
}
