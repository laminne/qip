package token

import (
	"crypto/rand"
	"fmt"
)

func secureRandom(keyLength int) string {
	key := make([]byte, keyLength)
	if _, err := rand.Read(key); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", key)
}
