package utils

import (
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"
)

const length = 8

func hash(input string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(input))
	if err != nil {
		return []byte{}, err
	}
	return h.Sum(nil), err
}

func encode(bytes []byte, digits uint32) string {
	return base58.Encode(bytes)[:digits]
}

func EncodeUrl(url string) (string, error) {
	b, err := hash(url)
	if err != nil {
		return "", err
	}
	return encode(b, length), err
}
