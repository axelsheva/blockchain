package utils

import (
	"crypto/sha256"

	"github.com/jamesruan/sodium"
)

func GenerateKeyPair(secret string) *sodium.SignKP {
	h := sha256.New()

	_, err := h.Write([]byte(secret))
	if err != nil {
		panic(err)
	}

	seed := sodium.SignSeed{
		Bytes: h.Sum(nil),
	}
	signKP := sodium.SeedSignKP(seed)

	return &signKP
}
