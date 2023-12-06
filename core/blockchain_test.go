package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, bc.Height()+1)
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, uint32(lenBlocks), bc.Height())
	assert.Equal(t, len(bc.headers), lenBlocks+1)
}

func TestBlockchain(t *testing.T) {
	bc := newBlockWithGenesis(t)
	assert.Equal(t, uint32(0), bc.Height())
	assert.NotNil(t, bc.validator)
}

func TestHasBlock(t *testing.T) {
	bc := newBlockWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}
