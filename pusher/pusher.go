package pusher

import (
	"github.com/Blockchainpartner/claim.it-back/util"
	"github.com/pusher/pusher-http-go"
)

var PusherClient pusher.Client

func InitPusherClient() {
	// Initialize a Pusher client
	PusherClient = pusher.Client{
		AppId:   util.PusherAppId,
		Key:     util.PusherAppKey,
		Secret:  util.PusherAppSecret,
		Cluster: util.PusherAppCluster,
		Secure:  true,
	}
}
