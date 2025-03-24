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
	plaintext, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	ciphertext := shift.Encipher(plaintext, byte(*key))

	os.Stdout.Write(ciphertext)
}
