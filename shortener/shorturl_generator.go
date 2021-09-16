package shortener

import (
	"crypto/sha256"
	"fmt"
	"os"
	"math/big"
	"github.com/itchyny/base58-go"
)

func sha256Encrypt(input string) []byte{
	encrypt := sha256.New()
	encrypt.Write([]byte(input))
	return encrypt.Sum(nil)
}

func base58Encoded(theBytes []byte) string{
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(theBytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// generate short url with initialURL from base58 + userId
// to prevent similar shortened urls
func GenerateShortLink(initialURL string, userId string) string {
	urlHashBytes := sha256Encrypt(initialURL + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}