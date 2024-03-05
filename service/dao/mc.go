package dao

import (
	"foragerServer/options"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Memcache struct {
	client *memcache.Client
}

func NewMC(c *options.MCConfig) *Memcache {
	client := memcache.New(c.Addr)
	client.MaxIdleConns = c.Idle
	client.Timeout = time.Duration(c.Timeout) * time.Millisecond
	return &Memcache{
		client: client,
	}
}

func (mc *Memcache) Get(key string) (string, error) {
	item, err := mc.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (mc *Memcache) Set(key string, value string, expire int32) (err error) {
	item := &memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: expire,
	}
	if err = mc.client.Set(item); err != nil {
		return
	}
	return
}
