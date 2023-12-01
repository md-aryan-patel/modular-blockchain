package core

import (
	"crypto/sha256"

	"github.com/md-aryan-patel/projectx/types"
)

type Hasher[T any] interface {
	/* Generic Hash function and returns Hash */
	Hash(T) types.Hash
}
type BlockHasher struct{}

func (BlockHasher) Hash(b *Block) types.Hash {

	/* applying sha256 */
	h := sha256.Sum256(b.HeaderData())
	return types.Hash(h)
}
