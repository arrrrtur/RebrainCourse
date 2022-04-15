package functions

import (
	"crypto/rand"
	"math/big"
)

func City() string {
	cities := []string{"Moscow", "Bishkek", "London", "New-York", "Berlin"}

	r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(cities))))
	return cities[r.Int64()]
}
