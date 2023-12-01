package core

import (
	"fmt"

	"github.com/md-aryan-patel/projectx/crypto"
)

type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (t *Transaction) Sign(prevKey crypto.PrivateKey) error {
	sig, err := prevKey.Sign(t.Data)
	if err != nil {
		return err
	}

	t.PublicKey = prevKey.GetPublicKey()
	t.Signature = sig

	return nil
}

func (t *Transaction) Verify() error {
	if t.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !t.Signature.Verify(t.PublicKey, t.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
