package common

import (
	"crypto/rand"
	"math/big"
)

var rander = rand.Reader

var allowRune = []rune("qưertyuiopasdfghjklzxcvbnm0123456789QƯERTYUIOPASDFGHJKLZXCVBNMP")

func RuneRandom(lenght int) ([]rune, error) {
	keyAllow := big.NewInt(int64(len(allowRune)))

	result := make([]rune, lenght)

	for i := 0; i < lenght; i++ {
		key, err := rand.Int(rander, keyAllow)
		if err != nil {
			return result, err
		}

		result[i] = allowRune[key.Uint64()]

	}
	return result, nil
}

func MustString(length int) string {
	seq, err := RuneRandom(length)
	if err != nil {
		panic(err)
	}
	return string(seq)
}

func GenSalt() string {
	return MustString(15)
}
