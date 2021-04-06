package main

import (
	"context"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/multichain/chain/bitcoincash"
	"github.com/renproject/multichain/chain/digibyte"
	"github.com/renproject/multichain/chain/dogecoin"
	"github.com/renproject/multichain/api/utxo"

	// "github.com/renproject/multichain/chain/filecoin"
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
	client := bitcoin.NewClient(bitcoin.DefaultClientOptions())
	txBuilder := bitcoin.NewTxBuilder(&chaincfg.RegressionNetParams)

	tx, _ := txBuilder.BuildTx(nil, nil)

	ctx, _ := context.WithTimeout(context.Background(), 0 * time.Second)
	client.LatestBlock(ctx)
	client.Output(ctx, utxo.Outpoint{})
	client.UnspentOutput(ctx, utxo.Outpoint{})
	client.SubmitTx(ctx, tx)
	client.UnspentOutputs(ctx, 0, 0, "")
	client.Confirmations(ctx, []byte{})
	client.EstimateSmartFee(ctx, 0)
	client.EstimateFeeLegacy(ctx, 0)

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
	// filecoin.NewClient(filecoin.DefaultClientOptions())
	// filecoin.NewTxBuilder()
}

func useSolana() {}

func useSubstrate() {}

func useTerra() {}

func useZcash() {
	zcash.NewClient(zcash.DefaultClientOptions())
	zcash.NewTxBuilder(&zcash.RegressionNetParams, 1000000)
}
