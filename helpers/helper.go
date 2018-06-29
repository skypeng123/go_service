package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	_ "github.com/cihub/seelog"
	"github.com/pborman/uuid"
	"math"
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

func Timestamp() int64 {
	return time.Now().Unix()
}

func Intval(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return result
}

//日期时间转为时间戳
func TimeUnix(datetime string) int64 {
	loc, _ := time.LoadLocation("Local") //获取时区
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	return t.Unix()
}

//指定时间戳的明天的开始时间戳
func TomorrowUnix(timestamp int64) int64 {
	date_time := time.Unix(timestamp+86400, 0).Format("2006-01-02")
	return TimeUnix(date_time + " 00:00:00")
}

//将秒转为分钟
func CeilMin(sec int64) int64 {
	return int64(math.Ceil(float64(sec) / 60))
}

func FormatParkingFee(fee float64) float64 {
	result, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", fee), 64)
	return result
}
