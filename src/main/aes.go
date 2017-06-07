package main

import (
	"crypto/aes"
)

func main() {
	keyStr := "a14a3bae270d4034892879749d68caa5"
	aes.NewCipher([]byte(keyStr))
}
