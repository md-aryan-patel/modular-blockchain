package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/md-aryan-patel/projectx/crypto"
	"github.com/md-aryan-patel/projectx/types"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}

	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	key := crypto.GeneratePrivateKey()
	b := randomBlock(height)
	b.Sign(key)
	return b
}

func TestHashBlock(t *testing.T) {
	b := randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}
