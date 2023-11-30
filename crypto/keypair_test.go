package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey1 := GeneratePrivateKey()
	publicKey1 := privateKey1.GetPublicKey()
	address1 := publicKey1.Address()

	fmt.Println(address1)
	msg := []byte("Hello world!")
	sig, err := privateKey1.Sign(msg)
	assert.Nil(t, err)

	ver := sig.Verify(publicKey1, msg)
	assert.True(t, ver)
	fmt.Println(sig)
}
