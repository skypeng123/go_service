package services

import (
	"github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"github.com/json-iterator/go"
	"go_service/helpers"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func Handle404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "file not found"})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

//授权验证
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		seelog.Debugf("请求参数 token: %s", token)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 10401, "msg": "unauthorized"})
			c.Abort()
			return
		}

		//从连接池获取redis连接
		rc := helpers.RedisClient.Get()
		defer rc.Close()

		cache_token_val, err := redis.Bytes(rc.Do("GET", "goservice:token:"+token))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 10402, "msg": "unauthorized or token may have expired"})
			c.Abort()
			return
		}
		appid := jsoniter.Get(cache_token_val, "appid").ToString()
		cache_appid_val, err := redis.String(rc.Do("GET", "goservice:appid:"+appid))
		if err != nil || cache_appid_val != token {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 10402, "msg": "unauthorized or token may have expired"})
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
