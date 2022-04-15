package functions

import (
	"crypto/rand"
	"math/big"
)

func Digit() string {
	digits := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
	return digits[r.Int64()]
}
