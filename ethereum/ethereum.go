package ethereum

import (
	"github.com/Blockchainpartner/claim.it-back/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	Client     *ethclient.Client
	WsClient   *ethclient.Client
	Transactor *bind.TransactOpts
)

func EthClientInit() {
	// connecting to node via RPC
	var err error
	Client, err = ethclient.Dial(util.EthNodeUri)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// connecting to node via WS
	WsClient, err = ethclient.Dial(util.EthNodeWSUri)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// use the Private key to unlock the wallet
	ethPrivateKey, err := crypto.HexToECDSA(util.HexEthPrivateKey)
	if err != nil {
		log.Fatalf("failed to parse Ethereum private key: %v\n", err)
	}
	Transactor = bind.NewKeyedTransactor(ethPrivateKey)
	// force GasLimit
	Transactor.GasLimit = 1e6
}
