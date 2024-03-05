package dao

import (
	"foragerServer/options"

	"gorm.io/gorm"
)

type Dao struct {
	Db       *gorm.DB
	Redis    *Redis
	Memcache *Memcache
}

func NewDao(cfg *options.DaoConfig) *Dao {
	redis := NewRedis(cfg.Redis)
	mc := NewMC(cfg.Memcache)
	db := NewDB(cfg.Mysql)
	d := &Dao{
		Db:       db,
		Redis:    redis,
		Memcache: mc,
	}
	return d
}

func (dao *Dao) Close() {
	dao.Redis.Close()
}
