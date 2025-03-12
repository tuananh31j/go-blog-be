package hashser

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Hash(pwd, salt string) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%v%v", pwd, salt)))
	return hex.EncodeToString(hasher.Sum(nil))
}
