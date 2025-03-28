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
	flag.Parse()
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	plaintext, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	ciphertext := shift.Encipher(plaintext, key)

	os.Stdout.Write(ciphertext)
}
