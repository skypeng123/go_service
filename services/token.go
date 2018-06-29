package services

import (
	"github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"github.com/json-iterator/go"
	"go_service/helpers"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type TokenInfo struct {
	Token      string `json:"token"`
	Expires_in int    `json:"expires_in"`
}

type tokenCache struct {
	Appid string `json:"appid"`
	Time  int64  `json:"time"`
}

var AppAccounts = map[string]string{
	"33efac4290daa8607cc5541c73b9a597": "2c4518ae7270e7f5cd5b255a1f46e93b",
}

func Token(c *gin.Context) {
	//get request params
	appid := c.PostForm("appid")
	secret := c.PostForm("secret")

	seelog.Debugf("请求参数 appid: %s,secret:%s", appid, secret)

	if appid == "" || secret == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "unauthorized"})
		return
	}

	if AppAccounts[appid] != secret {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "unauthorized"})
		return
	}

	//从连接池获取redis连接
	rc := helpers.RedisClient.Get()
	defer rc.Close()

	token, err := redis.String(rc.Do("GET", "goservice:appid:"+appid))
	if token == "" || err != nil {
		token = helpers.Token()
	}

	timeout := 7200
	cache_data := tokenCache{appid, helpers.Timestamp()}
	encoded, _ := jsoniter.Marshal(cache_data)
	//保存至redis
	rc.Do("SETEX", "goservice:token:"+token, timeout, encoded)
	rc.Do("SETEX", "goservice:appid:"+appid, timeout, token)

	data := TokenInfo{token, timeout}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": data})
	return
}
