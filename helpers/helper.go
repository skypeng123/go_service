package helpers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pborman/uuid"
	"log"
	"strconv"
	"time"
)

var (
	SqlDB       *sql.DB
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_PWD   string
	REDIS_DB    int
)

func init() {
	initDb()
	initRedis()
}

func initDb() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:alk3306%@tcp(119.23.35.15:3306)/alkparking_mycat_db1?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer SqlDB.Close()

	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func initRedis() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = "127.0.0.1:6379"
	REDIS_PWD = ""
	REDIS_DB = 0
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     10,
		MaxActive:   20,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}

			if REDIS_PWD != "" {
				if _, err := c.Do("AUTH", REDIS_PWD); err != nil {
					c.Close()
					log.Fatal(err.Error())
					return nil, err
				}
			}

			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

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
