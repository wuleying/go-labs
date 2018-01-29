package commands

import (
	"github.com/go-clog/clog"
	"go-labs/silver-blockchain/src/server"
	"go-labs/silver-blockchain/src/util"
	"go-labs/silver-blockchain/src/wallet"
)

func startNode(nodeId string, minerAddress string) {
	clog.Info("Starting node %s", nodeId)

	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			clog.Info("Mining is on. Address to receive rewards: %s", minerAddress)
		} else {
			clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, "Wrong miner address.")
		}
	}

	server.StartServer(nodeId, minerAddress)
}
