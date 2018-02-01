package commands

import (
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-blockchain/src/block"
)

func reindexUTXO(nodeId string) {
	bc := block.NewBlockChain(nodeId)
	UTXOSet := block.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()

	clog.Info("Reindex UTXO done. There are %d transactions in the UTXO set.", count)
}
