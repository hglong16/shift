package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hglong16/shift"
)

type testCase struct {
	key                   []byte
	plaintext, ciphertext []byte
}

var cases = []testCase{
	{key: []byte{1}, plaintext: []byte("HAL"), ciphertext: []byte("IBM")},
	{key: []byte{1}, plaintext: []byte("ADD"), ciphertext: []byte("BEE")},
	{key: []byte{1}, plaintext: []byte("ANA"), ciphertext: []byte("BOB")},
	{key: []byte{1}, plaintext: []byte("INKS"), ciphertext: []byte("JOLT")},
	{key: []byte{1}, plaintext: []byte("ADMIX"), ciphertext: []byte("BENJY")},
	{key: []byte{1}, plaintext: []byte{0, 1, 2, 3, 255}, ciphertext: []byte{1, 2, 3, 4, 0}},
	{key: []byte{2}, plaintext: []byte("SPEC"), ciphertext: []byte("URGE")},
	{key: []byte{3}, plaintext: []byte("PERK"), ciphertext: []byte("SHUN")},
	{key: []byte{4}, plaintext: []byte("GEL"), ciphertext: []byte("KIP")},
	{key: []byte{7}, plaintext: []byte("CHEER"), ciphertext: []byte("JOLLY")},
	{key: []byte{10}, plaintext: []byte("BEEF"), ciphertext: []byte("LOOP")},
	{key: []byte{1, 2, 3}, plaintext: []byte("ABC"), ciphertext: []byte("BDF")},
	{key: []byte{1, 2, 3}, plaintext: []byte{0, 0, 0}, ciphertext: []byte{1, 2, 3}},
	{key: []byte{1, 2, 3}, plaintext: []byte{4, 5, 6}, ciphertext: []byte{5, 7, 9}},
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

func TestDecipherWithKey(t *testing.T) {
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

			if !bytes.Equal(tc.key, got) {
				t.Fatalf("got %d, want %d", got, tc.key)
			}

		})
	}
}
