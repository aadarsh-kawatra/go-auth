package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	keyLength  = 32       // length of derived key (e.g., for AES-256)
	iterations = 1        // number of iterations
	memory     = 64 << 10 // memory cost in KiB (~64MB)
	threads    = 4        // number of parallel threads
)

func genSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	return salt, err
}

func GenerateHash(value string) (string, error) {
	salt, err := genSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(value), []byte(salt), iterations, memory, threads, keyLength)

	encoded := fmt.Sprintf("%s:%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encoded, nil
}

func VerifyHash(value, encodedHash string) bool {
	parts := strings.Split(encodedHash, ":")
	if len(parts) != 2 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	// Recompute hash with same parameters
	hash := argon2.IDKey([]byte(value), salt, 1, 64*1024, 4, uint32(len(expectedHash)))

	// Constant time compare
	return subtle.ConstantTimeCompare(hash, expectedHash) == 1
}
