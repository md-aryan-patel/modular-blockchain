package types

import (
	"encoding/hex"
	"fmt"
)

/* type Address if of slice of uint8 of size 32 */
type Address [20]uint8

/* converts given address into slice of bytes */
func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

/*
Overriding string function

whenever we print address using fmt.println(), GO will automatically
convert the bytes hash to string
*/
func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

/* converting bytes to Address datatype */
func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("given bytes wiht lenght with %d should be 20", len(b))
		panic(msg)
	}

	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}

	return Address(value)
}
