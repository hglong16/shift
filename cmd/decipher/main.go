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
	keyHex := flag.String("key", "01", "key in hẽadecimal (for example 'FF')")
	key, err := hex.DecodeString(*keyHex)
	flag.Parse()
	ciphertext, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	plaintext := shift.Decipher(ciphertext, key)

	os.Stdout.Write(plaintext)
}
