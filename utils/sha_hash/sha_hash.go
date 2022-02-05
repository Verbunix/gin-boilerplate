package sha_hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
)

func CreateShaHash(str string) string {
	key := []byte(str)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(os.Getenv("PASSWORD_SALT")))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
