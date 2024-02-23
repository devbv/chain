package types

import (
	"math/big"
)

type BlockEventData struct {
	BlockNumber    *big.Int
	BlockHash      []byte
	Timestamp      uint64
	NftTransferred []*OwnershipTransferEvent
}

type OwnershipTransferEvent struct {
	TxId         string
	BlockNumber  uint64
	ContractAddr []byte
	To           []byte
	TokenId      *big.Int
}
