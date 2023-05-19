package util

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"newsletter-platform/service/errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
}

// Function for creating a hash from a string and salt.
func HashPassword(password, salt string) (hash string) {
	hashFunctionParams := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		keyLength:   32,
	}
	saltBytes := []byte(salt)

	passwordHash := argon2Hash([]byte(password), saltBytes, hashFunctionParams)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		hashFunctionParams.memory,
		hashFunctionParams.iterations,
		hashFunctionParams.parallelism,
		base64.RawStdEncoding.EncodeToString(saltBytes),
		base64.RawStdEncoding.EncodeToString(passwordHash),
	)
}

// Function for validating if a password is comparable to a hash.
func ComparePasswordHash(password, hash string) (match bool, err error) {
	parameters, salt, outputHash, err := decodeHash(hash)
	if err != nil {
		return
	}
	passwordHash := argon2Hash([]byte(password), salt, parameters)

	match = subtle.ConstantTimeCompare(passwordHash, outputHash) == 1
	if subtle.ConstantTimeCompare(passwordHash, outputHash) == 1 {
		return true, nil
	}
	return
}

// Function for using an argon2 algorithm for password hashing.
func argon2Hash(password, salt []byte, hashParams *params) []byte {
	return argon2.IDKey(
		[]byte(password),
		salt,
		hashParams.iterations,
		hashParams.memory,
		hashParams.parallelism,
		hashParams.keyLength,
	)
}

// Function for decoding values in a hash.
func decodeHash(input string) (p *params, salt []byte, hash []byte, err error) {
	values := strings.Split(input, "$")
	if len(values) != 6 {
		return nil, nil, nil, errors.ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))
	return p, salt, hash, nil
}
