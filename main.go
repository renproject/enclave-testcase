package main

import (
	"context"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/renproject/multichain"
	"github.com/renproject/multichain/api/utxo"
	"github.com/renproject/multichain/chain/bitcoin"
	"github.com/renproject/multichain/chain/bitcoincash"
	// "github.com/renproject/multichain/chain/cosmos"
	"github.com/renproject/multichain/chain/digibyte"
	"github.com/renproject/multichain/chain/dogecoin"
	"github.com/renproject/multichain/chain/filecoin"
	"github.com/renproject/multichain/chain/solana"
	"github.com/renproject/multichain/chain/terra"
	"github.com/renproject/multichain/chain/zcash"
	"github.com/renproject/enclave-testcase/ethereumbinding"
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

	ctx, _ := context.WithTimeout(context.Background(), 0*time.Second)
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

func useEthereum() {
	client, err := ethclient.Dial(string(chainOpts.RPC))
	if err != nil {
		return
	}
	gatewayRegistry, err := ethereumbinding.NewGatewayRegistry(common.HexToAddress(""), client)
	if err != nil {
		return
	}
	gateways := make(map[multichain.Asset]*ethereumbinding.MintGatewayLogicV1, 0)
	assetAddrs := make(map[multichain.Address]multichain.Asset, 0)
	sourceChains := make(multichain.Chain, 0)
	for _, chain := range sourceChains {
		gatewayAddr, err := gatewayRegistry.GetGatewayBySymbol(&bind.CallOpts{}, string(chain.NativeAsset()))
		if err != nil {
			return
		}
		gateway, err := ethereumbinding.NewMintGatewayLogicV1(gatewayAddr, client)
		if err != nil {
			return
		}
		gateways[chain.NativeAsset()] = gateway

		addr, err := gateway.Token(&bind.CallOpts{})
		if err != nil {
			return
		}
		assetAddrs[multichain.Address(addr.String())] = chain.NativeAsset()
	}
}

func useFilecoin() {
	filecoin.NewClient(filecoin.DefaultClientOptions())
	filecoin.NewTxBuilder()
}

func useSolana() {
	solana.NewClient(solana.DefaultClientOptions())
}

func useSubstrate() {}

func useTerra() {
	client := terra.NewClient(terra.DefaultClientOptions())
	terra.NewTxBuilder(terra.DefaultTxBuilderOptions(), client)
}

func useZcash() {
	zcash.NewClient(zcash.DefaultClientOptions())
	zcash.NewTxBuilder(&zcash.RegressionNetParams, 1000000)
}
