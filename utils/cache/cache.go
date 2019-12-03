package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var c *cache.Cache

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func Get(key string) (interface{}, bool) {
	// 读取缓存
	vaule, is := c.Get(key)
	return vaule, is
}

func Set(key string, vaule interface{}) {
	// 写入缓存 默认过期时间
	c.Set(key, vaule, cache.DefaultExpiration)
}

func Setx(key string, vaule interface{}) {
	// 写入缓存 永不过期
	c.Set(key, vaule, cache.NoExpiration)
}
