package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hglong16/shift"
)

func main() {
	crib := flag.String("crib", "", "crib text")

	flag.Parse()

	if *crib == "" {
		fmt.Println("crib is required")
		os.Exit(1)
	}

	ciphertext, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	key, err := shift.Crack(ciphertext, []byte(*crib))
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	block, err := shift.NewCipher(key)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	mode := shift.NewDecrypter(block)
	plaintext := make([]byte, len(ciphertext))

	mode.CryptBlocks(plaintext, ciphertext)
	plaintext = shift.Unpad(plaintext, mode.BlockSize())
	os.Stdout.Write(plaintext)

}
