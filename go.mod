module github.com/renproject/enclave-testcase

go 1.16

require (
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/ethereum/go-ethereum v1.9.20
	github.com/renproject/multichain v0.3.4
)

replace github.com/CosmWasm/go-cosmwasm => github.com/terra-project/go-cosmwasm v0.10.1-terra

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

replace github.com/renproject/solana-ffi => ./extern/solana-ffi
