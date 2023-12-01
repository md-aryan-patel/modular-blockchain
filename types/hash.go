package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

/* type Hash is of slice of uint8 of size 32 */
type Hash [32]uint8

/* checks if the hash provided is zero hash or not */
func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

/* converts given hash into slice of bytes */
func (h Hash) ToSlice() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = h[i]
	}
	return b
}

/*
Overriding string function

whenever we print hash using fmt.println(), GO will automatically
convert the bytes hash to string
*/

func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

/* converting bytes to Hash datatype */
func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("given bytes wiht lenght with %d should be 32", len(b))
		panic(msg)
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

/* generating random bytes */
func RandomByte(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

/* generates random hash */
func RandomHash() Hash {
	return HashFromBytes(RandomByte(32))
}
