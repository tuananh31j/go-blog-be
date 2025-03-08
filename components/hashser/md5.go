package hashser

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type md5Hash struct {
	password string
	salt     string
}

func NewMd5Hash(pass string) *md5Hash {
	return &md5Hash{
		password: pass,
	}
}

func (h *md5Hash) SetSalt(s string) {
	h.salt = s
}

func (h *md5Hash) Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%v%v", h.password, h.salt)))
	return hex.EncodeToString(hasher.Sum(nil))
}
