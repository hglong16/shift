package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hglong16/shift"
)

func main() {
	key := flag.Int("key", 1, "key to shift by")
	flag.Parse()
	ciphertext, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	plaintext := shift.Decipher(ciphertext, byte(*key))

	os.Stdout.Write(plaintext)
}
