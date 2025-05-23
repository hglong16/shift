package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hglong16/shift"
)

func main() {
	keyHex := flag.String("key", "", "32-byte key in hexadecimal")
	flag.Parse()
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	block, err := shift.NewCipher(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	enc := shift.NewEncrypter(block)
	plaintext, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext = shift.Pad(plaintext, enc.BlockSize())

	ciphertext := make([]byte, len(plaintext))

	enc.CryptBlocks(ciphertext, plaintext)
	os.Stdout.Write(ciphertext)
}
