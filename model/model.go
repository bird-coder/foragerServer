package model

import (
	"foragerServer/service/dao"
	"reflect"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	daoModel *DaoModel
)

var (
	ErrNoRecord = errors.New("record not found")
	ErrDatabase = errors.New("sql error")
)

func wrapError(res *gorm.DB, query string) (err error) {
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// err = errors.Wrapf(ErrNoRecord, "query: %s", query)
	} else {
		err = errors.Wrapf(ErrDatabase, "query: %s; error(%v)", query, res.Error)
	}
	return
}

type DaoModel struct {
	*dao.Dao
}

func Init(dao *dao.Dao) {
	daoModel = &DaoModel{
		dao,
	}
}

type ModelCheck interface {
	IsEmpty() bool
}

func (user *User) IsEmpty() bool {
	return reflect.DeepEqual(user, &User{})
}

func (address *UserAddress) IsEmpty() bool {
	return reflect.DeepEqual(address, &UserAddress{})
}

func (shop *Shop) IsEmpty() bool {
	return reflect.DeepEqual(shop, &Shop{})
}
