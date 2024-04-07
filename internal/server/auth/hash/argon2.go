package hash

import (
	"bytes"
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
)

// Error messages
var (
	ErrHashMismatch = errors.New("hash mismatch")
)

// Argon2IDConfig is a configuration for Argon2IDConfig hashing
type Argon2IDConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
	saltLen uint32
}

type Argon2ID struct {
	Config *Argon2IDConfig
}

var _ HashIface = (*Argon2ID)(nil)

// NewArgon2IDWithConfig creates a new Argon2ID instance
func NewArgon2IDWithConfig(config *Argon2IDConfig) *Argon2ID {
	return &Argon2ID{
		Config: config,
	}
}

// NewArgon2ID creates a new Argon2ID instance with default values
func NewArgon2ID() *Argon2ID {
	config := Argon2IDConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
		saltLen: 16,
	}

	return NewArgon2IDWithConfig(&config)
}

// GenerateSalt generates a random salt
func (a *Argon2ID) GenerateSalt() (*[]byte, error) {
	salt := make([]byte, a.Config.saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return &salt, nil
}

// GenerateHash generates a hash
func (a *Argon2ID) GenerateHash(data *[]byte, salt *[]byte) (*HashSalt, error) {
	hash := argon2.IDKey(*data, *salt, a.Config.time, a.Config.memory, a.Config.threads, a.Config.keyLen)
	hashSalt := &HashSalt{
		Hash: &hash,
		Salt: salt,
	}

	return hashSalt, nil
}

// GenerateSaltAndHash generates a salt and hash
func (a *Argon2ID) GenerateSaltAndHash(data *[]byte) (*HashSalt, error) {
	salt, err := a.GenerateSalt()
	if err != nil {
		return nil, err
	}

	return a.GenerateHash(data, salt)
}

// CompareHash compares a hash with data and salt
func (a *Argon2ID) CompareHash(data *[]byte, hash HashSalt) error {
	newHash, err := a.GenerateHash(data, hash.Salt)
	if err != nil {
		return err
	}

	if !bytes.Equal(*hash.Hash, *newHash.Hash) {
		return ErrHashMismatch
	}

	return nil
}
