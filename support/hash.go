package support

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data []byte) string {
	hash := md5.New()
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}
