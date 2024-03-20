package random

import (
	"crypto/rand"
	"encoding/hex"
)

func String(length int) string {
	tmp := make([]byte, length)
	_, _ = rand.Read(tmp)
	return hex.EncodeToString(tmp)
}
