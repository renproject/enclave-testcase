package main

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/multichain/chain/bitcoincash"
	"github.com/renproject/multichain/chain/digibyte"
	"github.com/renproject/multichain/chain/dogecoin"
	"github.com/renproject/multichain/chain/filecoin"
	"github.com/renproject/multichain/chain/zcash"
)

func main() {
	useBitcoin()
	useBitcoinCash()
	// useCosmos()
	useDigibyte()
	useDogecoin()
	useEthereum()
	useFilecoin()
	useSolana()
	useSubstrate()
	useTerra()
	useZcash()
}

func useBitcoin() {
	bitcoin.NewClient(bitcoin.DefaultClientOptions())
	bitcoin.NewTxBuilder(&chaincfg.RegressionNetParams)
}

func useBitcoinCash() {
	bitcoincash.NewClient(bitcoincash.DefaultClientOptions())
	bitcoincash.NewTxBuilder(&chaincfg.RegressionNetParams)
}

// func useCosmos() {}

func useDigibyte() {
	digibyte.NewClient(digibyte.DefaultClientOptions())
	digibyte.NewTxBuilder(&digibyte.RegressionNetParams)
}

func useDogecoin() {
	dogecoin.NewClient(dogecoin.DefaultClientOptions())
	dogecoin.NewTxBuilder(&dogecoin.RegressionNetParams)
}

func useEthereum() {}

func useFilecoin() {
	filecoin.NewClient(filecoin.DefaultClientOptions())
	filecoin.NewTxBuilder()
}

func useSolana() {}

func useSubstrate() {}

func useTerra() {}

func useZcash() {
	zcash.NewClient(zcash.DefaultClientOptions())
	zcash.NewTxBuilder(&zcash.RegressionNetParams, 1000000)
}
