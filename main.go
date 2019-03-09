package main

import (
	"github.com/Blockchainpartner/claim.it-back/db"
	"github.com/Blockchainpartner/claim.it-back/ethereum"
	"github.com/Blockchainpartner/claim.it-back/pusher"
	"github.com/Blockchainpartner/claim.it-back/server"
	"github.com/Blockchainpartner/claim.it-back/util"
)

func main() {
	// init environment variables
	util.InitEnvironment()
	// init db connection
	db.Init()
	// init Pusher client
	pusher.InitPusherClient()
	// init Ethereum client
	ethereum.EthClientInit()
	// init and start server
	server.Init()
}
