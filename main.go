package main

import (
	"fmt"
	"strings"
)

// JeffersonDiskCipher represents the Jefferson Disk Cipher.
type JeffersonDiskCipher struct {
	disk       []rune
	rotations  int
	plaintext  string
	ciphertext string
}

// NewJeffersonDiskCipher creates a new instance of the Jefferson Disk Cipher.
func NewJeffersonDiskCipher(rotations int) *JeffersonDiskCipher {
	if rotations < 0 || rotations >= 26 {
		panic("Invalid number of rotations. Rotations must be between 0 and 25.")
	}
	disk := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return &JeffersonDiskCipher{disk: disk, rotations: rotations}
}

// Encrypt encrypts the plaintext message using the Jefferson Disk Cipher.
func (j *JeffersonDiskCipher) Encrypt(plaintext string) string {
	plaintext = strings.ToUpper(plaintext)
	ciphertext := ""

	for _, char := range plaintext {
		if char == ' ' {
			ciphertext += " "
			continue
		}

		index := strings.IndexRune(string(j.disk), char)
		if index == -1 {
			ciphertext += string(char) // Preserve characters not in the disk
		} else {
			rotatedDisk := append(j.disk[j.rotations:], j.disk[:j.rotations]...)
			ciphertext += string(rotatedDisk[index])
		}
	}

	j.plaintext = plaintext
	j.ciphertext = ciphertext
	return ciphertext
}

// Decrypt decrypts the ciphertext message using the Jefferson Disk Cipher.
func (j *JeffersonDiskCipher) Decrypt(ciphertext string) string {
	decrypted := ""

	for _, char := range ciphertext {
		if char == ' ' {
			decrypted += " "
			continue
		}

		index := strings.IndexRune(string(j.disk), char)
		if index == -1 {
			decrypted += string(char) // Preserve characters not in the disk
		} else {
			rotatedDisk := append(j.disk[len(j.disk)-j.rotations:], j.disk[:len(j.disk)-j.rotations]...)
			decrypted += string(rotatedDisk[index])
		}
	}

	return decrypted
}

func main() {
	rotations := 3 // Example number of rotations
	cipher := NewJeffersonDiskCipher(rotations)

	plaintext := "HELLO WORLD"
	ciphertext := cipher.Encrypt(plaintext)
	decrypted := cipher.Decrypt(ciphertext)

	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Decrypted:", decrypted)
}
