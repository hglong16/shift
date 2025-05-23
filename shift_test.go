package shift_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/hglong16/shift"
)

var testKey = bytes.Repeat([]byte{1}, shift.BlockSize)

func TestNewCipeher_GivesErrorForInvalidKey(t *testing.T) {
	t.Parallel()

	_, err := shift.NewCipher([]byte{})

	if !errors.Is(err, shift.ErrKeySize) {
		t.Fatalf("NewCipher() = %v, want %v", err, shift.ErrKeySize)
	}

}

func TestNewCiper_GivesNoErrorForValidKey(t *testing.T) {
	t.Parallel()

	_, err := shift.NewCipher(make([]byte, shift.BlockSize))
	if err != nil {
		t.Fatalf("NewCipher() = %v, want nil", err)
	}
}

func TestBlockSize_ReturnBlockSize(t *testing.T) {
	t.Parallel()

	block, err := shift.NewCipher(make([]byte, shift.BlockSize))
	if err != nil {
		t.Fatal(err)
	}
	want := shift.BlockSize
	got := block.BlockSize()

	if want != got {
		t.Errorf("BlockSize() = %d, want %d", got, want)
	}

}

var cipherCases = []struct {
	plaintext, ciphertext []byte
}{
	{
		plaintext:  []byte{0, 1, 2, 3, 4, 5},
		ciphertext: []byte{1, 2, 3, 4, 5, 6},
	},
}

func TestEncrypt(t *testing.T) {
	t.Parallel()

	block, err := shift.NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cipherCases {
		name := fmt.Sprintf("%x + %x = %x", tc.plaintext, testKey, tc.ciphertext)
		t.Run(name, func(t *testing.T) {
			got := make([]byte, len(tc.plaintext))
			block.Encrypt(got, tc.plaintext)
			if !bytes.Equal(tc.ciphertext, got) {
				t.Errorf("Encrypt() = %x, want %x", got, tc.ciphertext)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	t.Parallel()

	block, err := shift.NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cipherCases {
		name := fmt.Sprintf("%x + %x = %x", tc.plaintext, testKey, tc.ciphertext)
		t.Run(name, func(t *testing.T) {
			got := make([]byte, len(tc.ciphertext))
			block.Decrypt(got, tc.ciphertext)
			if !bytes.Equal(tc.plaintext, got) {
				t.Errorf("Encrypt() = %x, want %x", got, tc.ciphertext)
			}
		})
	}
}

func TestEncryptEnciphersBlockAlignedMessage(t *testing.T) {
	t.Parallel()

	plaintext := []byte("This message is exactly 32 bytes")

	block, err := shift.NewCipher(testKey)

	if err != nil {
		t.Fatal(err)
	}

	enc := shift.NewEncrypter(block)
	want := []byte("Uijt!nfttbhf!jt!fybdumz!43!czuft")
	got := make([]byte, 32)
	enc.CryptBlocks(got, plaintext)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

}

func TestEncrypterCorrectlyReportsCipherBlockSize(t *testing.T) {
	t.Parallel()

	block, err := shift.NewCipher(testKey)

	if err != nil {
		t.Fatal(err)
	}

	enc := shift.NewEncrypter(block)
	want := shift.BlockSize
	got := enc.BlockSize()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestDecrypterEnciphersBlockAlignedMessage(t *testing.T) {
	t.Parallel()

	ciphertext := []byte("Uijt!nfttbhf!jt!fybdumz!43!czuft")

	block, err := shift.NewCipher(testKey)

	if err != nil {
		t.Fatal(err)
	}

	enc := shift.NewDecrypter(block)
	want := []byte("This message is exactly 32 bytes")
	got := make([]byte, 32)
	enc.CryptBlocks(got, ciphertext)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

}

func TestDecrypterCorrectlyReportsCipherBlockSize(t *testing.T) {
	t.Parallel()

	block, err := shift.NewCipher(testKey)

	if err != nil {
		t.Fatal(err)
	}

	enc := shift.NewDecrypter(block)
	want := shift.BlockSize
	got := enc.BlockSize()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

var padCases = []struct {
	name        string
	raw, padded []byte
}{
	{
		name:   "1 short of full block",
		raw:    []byte{0, 0, 0},
		padded: []byte{0, 0, 0, 1},
	},
	{
		name:   "2 short of full block",
		raw:    []byte{0, 0},
		padded: []byte{0, 0, 2, 2},
	},
	{
		name:   "3 short of full block",
		raw:    []byte{0},
		padded: []byte{0, 3, 3, 3},
	},
	{
		name:   "full block",
		raw:    []byte{0, 0, 0, 0},
		padded: []byte{0, 0, 0, 0, 4, 4, 4, 4},
	},
	{
		name:   "empty block",
		raw:    []byte{},
		padded: []byte{4, 4, 4, 4},
	},
}

func TestPad(t *testing.T) {
	t.Parallel()

	blockSize := 4

	for _, tc := range padCases {
		t.Run(tc.name, func(t *testing.T) {
			got := shift.Pad(tc.raw, blockSize)
			if !bytes.Equal(tc.padded, got) {
				t.Errorf("Pad() = %x, want %x", got, tc.padded)
			}
		})
	}
}

func TestUnpad(t *testing.T) {
	t.Parallel()

	blockSize := 4

	for _, tc := range padCases {
		t.Run(tc.name, func(t *testing.T) {
			got := shift.Unpad(tc.padded, blockSize)
			// if err != nil {
			// 	t.Fatal(err)
			// }
			if !bytes.Equal(tc.raw, got) {
				t.Errorf("Unpad() = %x, want %x", got, tc.raw)
			}
		})
	}
}

func TestCrack(t *testing.T) {
	t.Parallel()
	plaintext := []byte("This message is exactly 32 bytes")
	ciphertext := []byte("Uijs message is exactly 32 bytes")
	want := append([]byte{1, 1, 1}, bytes.Repeat([]byte{0}, 29)...)
	got, err := shift.Crack(ciphertext, plaintext)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(want, got) {
		t.Fatalf("want %x, got %x", want, got)
	}
}
func TestNextCorrectlyIncrementsInputWithoutOverflow(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		input, want []byte
	}{
		{
			input: []byte{0, 0, 0},
			want:  []byte{1, 0, 0},
		},
		{
			input: []byte{255, 0, 0},
			want:  []byte{0, 1, 0},
		},
		{
			input: []byte{255, 255, 0},
			want:  []byte{0, 0, 1},
		},
		{
			input: []byte{255, 255, 254},
			want:  []byte{0, 0, 255},
		},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%x", tc.input), func(t *testing.T) {
			got, err := shift.Next(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(tc.want, got) {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestNextReturnsErrorWhenNextKeyWouldOverflow(t *testing.T) {
	t.Parallel()
	_, err := shift.Next([]byte{255, 255, 255})
	if err == nil {
		t.Fatal("want error on key overflow, but got nil")
	}
}

func BenchmarkCrack(b *testing.B) {
	plaintext := []byte("This message is exactly 32 bytes")
	ciphertext := []byte("Uiis message is exactly 32 bytes")
	b.ResetTimer()
	for range b.N {
		_, _ = shift.Crack(ciphertext, plaintext)
	}
}
