package core

import (
	"testing"

	"github.com/md-aryan-patel/projectx/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	tx := &Transaction{
		Data: []byte("foo"),
	}
	prevKey := crypto.GeneratePrivateKey()
	err := tx.Sign(prevKey)

	assert.Nil(t, err)
	assert.Equal(t, prevKey.GetPublicKey().Address(), tx.PublicKey.Address())
}

func TestVerifyTransaction(t *testing.T) {
	tx := &Transaction{
		Data: []byte("foo"),
	}
	prevKey := crypto.GeneratePrivateKey()
	err := tx.Sign(prevKey)
	assert.Nil(t, err)

	assert.Nil(t, tx.Verify())
}
