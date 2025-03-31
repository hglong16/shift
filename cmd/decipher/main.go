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
	keyHex := flag.String("key", "", "32-bytes key in hexadecimal")
	flag.Parse()
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	block, err := shift.NewCipher(key)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	ciphertext, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	plaintext := make([]byte, len(ciphertext))
	dec := shift.NewDecrypter(block)
	dec.CryptBlocks(plaintext, ciphertext)
	plaintext = shift.Unpad(plaintext, shift.BlockSize)
	os.Stdout.Write(plaintext)
}
