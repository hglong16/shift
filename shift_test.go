package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hglong16/shift"
)

type testCase struct {
	key                   byte
	plaintext, ciphertext []byte
}

var cases = []testCase{
	{key: 1, plaintext: []byte("HAL"), ciphertext: []byte("IBM")},
	{key: 1, plaintext: []byte("ADD"), ciphertext: []byte("BEE")},
	{key: 1, plaintext: []byte("ANA"), ciphertext: []byte("BOB")},
	{key: 1, plaintext: []byte("INKS"), ciphertext: []byte("JOLT")},
	{key: 1, plaintext: []byte("ADMIX"), ciphertext: []byte("BENJY")},
	{key: 1, plaintext: []byte{0, 1, 2, 3, 255}, ciphertext: []byte{1, 2, 3, 4, 0}},
	{key: 2, plaintext: []byte("SPEC"), ciphertext: []byte("URGE")},
	{key: 3, plaintext: []byte("PERK"), ciphertext: []byte("SHUN")},
	{key: 4, plaintext: []byte("GEL"), ciphertext: []byte("KIP")},
	{key: 7, plaintext: []byte("CHEER"), ciphertext: []byte("JOLLY")},
	{key: 10, plaintext: []byte("BEEF"), ciphertext: []byte("LOOP")},
}

func TestEncipherTransform(t *testing.T) {
	t.Parallel()
	for _, tc := range cases {
		name := fmt.Sprintf("%s + %d = %s", tc.plaintext, tc.key, tc.ciphertext)
		t.Run(name, func(t *testing.T) {
			got := shift.Encipher(tc.plaintext, tc.key)
			if !bytes.Equal(got, tc.ciphertext) {
				t.Errorf("got %q, want %q", got, tc.ciphertext)
			}
		})
	}

}

func TestDecipherWithKey1TransformsIBMToHAL(t *testing.T) {
	t.Parallel()
	for _, tc := range cases {
		name := fmt.Sprintf("%s - %d = %s", tc.ciphertext, tc.key, tc.plaintext)
		t.Run(name, func(t *testing.T) {
			got := shift.Decipher(tc.ciphertext, tc.key)
			if !bytes.Equal(got, tc.plaintext) {
				t.Errorf("got %q, want %q", got, tc.plaintext)
			}
		})
	}
}

func TestCrack(t *testing.T) {
	t.Parallel()

	for _, tc := range cases {
		name := fmt.Sprintf("%s + %d = %s", tc.plaintext, tc.key, tc.ciphertext)

		t.Run(name, func(t *testing.T) {
			got, err := shift.Crack(tc.ciphertext, tc.plaintext[:3])
			if err != nil {
				t.Fatal(err)
			}

			if tc.key != got {
				t.Fatalf("got %d, want %d", got, tc.key)
			}

		})
	}
}

func TestCrackReturnsErrorWhenKeyNotFound(t *testing.T) {
	t.Parallel()

	_, err := shift.Crack([]byte("no     good"), []byte("bogus"))
	if err == nil {
		t.Fatal("expected an error, but got nil")
	}
}
