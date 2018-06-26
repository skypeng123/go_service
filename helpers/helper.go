package helpers

import (
	"crypto/md5"
	"encoding/hex"
	_ "github.com/cihub/seelog"
	"github.com/pborman/uuid"
	"strconv"
	"time"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func UUID() string {
	return uuid.New()
}

func Token() string {
	return Md5(UUID())
}

func Timestamp() int {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	ts, _ := strconv.Atoi(timestamp[:10])
	return ts
}
