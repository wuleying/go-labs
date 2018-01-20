package commands

import (
	"github.com/go-clog/clog"
	"go-labs/silver-blockchain/src/block"
)

func reindexUTXO(address string) {
	bc := block.NewBlockChain(address)
	UTXOSet := block.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()

	clog.Info("Reindex UTXO done. There are %d transactions in the UTXO set.", count)
}
