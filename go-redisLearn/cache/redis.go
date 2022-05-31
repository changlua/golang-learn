package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	. "go-redisLearn/config" //这种方式引用，就不需要带上包前缀
	"time"
)

var ctx = context.Background()
var rc *redis.Client

func init() {
	host := AppConfig.Redis.Host
	port := AppConfig.Redis.Port
	password := AppConfig.Redis.Password
	rc = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}

func RCGet(key string) (string, error) {
	return rc.Get(ctx, key).Result()
}
func RCExists(key string) bool {
	return rc.Exists(ctx, key).Val() != 0
}

func RCSet(key string, value interface{}, expiration time.Duration) {
	if RCExists(key) {
		rc.Expire(ctx, key, expiration)
		return
	}
	rc.Set(ctx, key, value, expiration)
}

func RCIncrement(key string) {
	rc.Incr(ctx, key)
}
