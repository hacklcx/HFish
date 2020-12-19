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
	// Read cache
	vaule, is := c.Get(key)
	return vaule, is
}

func Set(key string, vaule interface{}) {
	// Write cache default expiration time
	c.Set(key, vaule, cache.DefaultExpiration)
}

func Setx(key string, vaule interface{}) {
	// Write to the cache and never expire
	c.Set(key, vaule, cache.NoExpiration)
}
