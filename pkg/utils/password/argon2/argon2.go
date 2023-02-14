package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/approvers/qip/pkg/utils/password"
	"golang.org/x/crypto/argon2"
)

const hashAlgorithm = "Argon2"

// createHashedPassword ハッシュ済みパスワードを返す
func (e *PasswordEncoder) createHashedPassword(rawPassword string, salt string) []byte {
	return argon2.IDKey([]byte(rawPassword), []byte(salt), 2, 20, 20, 20)
}

type PasswordEncoder struct{}

func NewArgon2PasswordEncoder() *PasswordEncoder {
	return &PasswordEncoder{}
}

func (e *PasswordEncoder) EncodePassword(rawPassword string) (password.EncodedPassword, error) {
	random, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return "", err
	}

	salt := fmt.Sprintf("%x[0]", random)
	hashedPassword := e.createHashedPassword(rawPassword, salt)
	// ハッシュ済みパスワードをhexにエンコード
	encodedHexedPassword := hex.EncodeToString(hashedPassword)
	//<hash algorithm>.<hashed Password as hex string>.<salt as hex string>
	combinatedAlgoAndHashAndPassword := fmt.Sprintf("%s.%s.%s", hashAlgorithm, encodedHexedPassword, salt)

	return password.EncodedPassword(combinatedAlgoAndHashAndPassword), nil
}

func (e *PasswordEncoder) IsMatchPassword(raw string, encoded password.EncodedPassword) bool {
	passwordPart, salt, err := e.decodeHash(encoded)
	if err != nil {
		return false
	}

	otherHash := e.createHashedPassword(raw, salt)

	return subtle.ConstantTimeCompare([]byte(passwordPart), otherHash) == 1
}

// decodeHash ハッシュ済みパスワード部、ソルトを返します
func (e *PasswordEncoder) decodeHash(encodedPassword password.EncodedPassword) (string, string, error) {
	split := strings.Split(string(encodedPassword), ".")
	if len(split) != 3 {
		return "", "", errors.New("input is invalid format")
	}
	encodedPasswordPart, salt := split[1], split[2]
	decodedPasswordPart, err := hex.DecodeString(encodedPasswordPart)
	if err != nil {
		return "", "", err
	}

	return string(decodedPasswordPart), salt, nil
}
