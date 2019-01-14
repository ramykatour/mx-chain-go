package block

import (
	"github.com/ElrondNetwork/elrond-go-sandbox/data/block"
	"github.com/ElrondNetwork/elrond-go-sandbox/data/transaction"
	"github.com/ElrondNetwork/elrond-go-sandbox/p2p"
	"github.com/ElrondNetwork/elrond-go-sandbox/process"
)

func (bp *blockProcessor) GetTransactionFromPool(destShardID uint32, txHash []byte) *transaction.Transaction {
	return bp.getTransactionFromPool(destShardID, txHash)
}

func (bp *blockProcessor) RequestTransactionFromNetwork(body *block.TxBlockBody) {
	bp.requestBlockTransactions(body)
}

func (bp *blockProcessor) WaitForTxHashes() {
	bp.waitForTxHashes()
}

func (bp *blockProcessor) ReceivedTransaction(txHash []byte) {
	bp.receivedTransaction(txHash)
}

func (hi *HeaderInterceptor) ProcessHdr(hdr p2p.Creator, rawData []byte) error {
	return hi.processHdr(hdr, rawData)
}

func (gbbi *GenericBlockBodyInterceptor) ProcessBodyBlock(bodyBlock p2p.Creator, rawData []byte) error {
	return gbbi.processBodyBlock(bodyBlock, rawData)
}

func (hdrRes *headerResolver) ResolveHdrRequest(rd process.RequestData) ([]byte, error) {
	return hdrRes.resolveHdrRequest(rd)
}

func (gbbRes *genericBlockBodyResolver) ResolveBlockBodyRequest(rd process.RequestData) ([]byte, error) {
	return gbbRes.resolveBlockBodyRequest(rd)
}