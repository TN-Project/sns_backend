package session

import (
	"encoding/json"
	"log/slog"
	"sns_backend/pkg/common/random"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {

	REDIS_HOST := "redis:6379"

	client = redis.NewClient(&redis.Options{
		Addr: REDIS_HOST,
		DB:   0,
	})

	if _, err := client.Ping(&gin.Context{}).Result(); err != nil {
		slog.Error("Error connecting to redis: " + err.Error())
		return
	}
}

type Session struct {
	SessionId string
	CookieKey string
	Model     interface{}
}

func Default(c *gin.Context, cookieKey string, model interface{}) *Session {
	SessionId, err := c.Cookie(cookieKey)
	if err != nil {
		SessionId := random.MakeRandomStringId(25)
		new(c, SessionId, cookieKey, model)
		return &Session{SessionId: SessionId, CookieKey: cookieKey, Model: model}
	}
	return &Session{SessionId: SessionId, CookieKey: cookieKey, Model: model}
}

func new(c *gin.Context, SessionId string, cookieKey string, value interface{}) {
	valueByte, err := json.Marshal(value)
	if err != nil {
		slog.Error("Error setting session: " + err.Error())
		return
	}
	client.Set(c, SessionId, string(valueByte), 24*30*time.Hour)
	c.SetCookie(cookieKey, SessionId, 0, "/", "", false, true)
}

func (s *Session) Set(c *gin.Context, value interface{}) {
	valueByte, err := json.Marshal(value)
	if err != nil {
		slog.Error("Error setting session: " + err.Error())
		return
	}
	client.Set(c, s.SessionId, string(valueByte), 24*30*time.Hour)
}

func (s *Session) Get(c *gin.Context) interface{} {
	SessionId, err := c.Cookie(s.CookieKey)
	if err != nil {
		return nil
	}

	value, err := client.Get(c, SessionId).Bytes()
	if err != nil {
		return nil
	}

	err = json.Unmarshal(value, s.Model)
	if err != nil {
		slog.Error("Error getting session: " + err.Error())
		return nil
	}

	return s.Model
}

func (s *Session) Delete(c *gin.Context) {
	client.Del(c, s.SessionId)
	c.SetCookie(s.CookieKey, "", -1, "/", "", false, true)
}
