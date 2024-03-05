package model

import (
	"fmt"

	"gorm.io/gorm"
)

func GetShopList(list []int) ([]Shop, error) {
	var shops []Shop
	var err error
	tx := daoModel.Db.Where(list)
	query := tx.ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Find(&shops) })
	if res := tx.Find(&shops); res.Error != nil {
		err = wrapError(res, query)
	}
	fmt.Printf("search shop list, ids: %v\n", list)
	return shops, err
}

func GetShopListByRange(minLat float64, maxLat float64, minLng float64, maxLng float64) ([]Shop, error) {
	var shops []Shop
	var err error
	tx := daoModel.Db.Where("lat > ? and lat < ? and lng > ? and lng < ?", minLat, maxLat, minLng, maxLng)
	query := tx.ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Find(&shops) })
	if res := tx.Find(&shops); res.Error != nil {
		err = wrapError(res, query)
	}
	fmt.Printf("search shop list by range, range: %f,%f,%f,%f\n", minLat, maxLat, minLng, maxLng)
	return shops, err
}

func GetShopInfo(id int) (Shop, error) {
	var shop Shop
	var err error
	tx := daoModel.Db.Where(&Shop{ID: id})
	query := tx.ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Find(&shop) })
	if res := tx.First(&shop); res.Error != nil {
		err = wrapError(res, query)
	}
	fmt.Printf("search shop info, id: %d\n", id)
	return shop, err
}

func GetShopGoods(shopId int) ([]Goods, error) {
	var goods []Goods
	var err error
	tx := daoModel.Db.Model(&Shop{}).Select("goods.*").
		Joins("inner join goods on shop.id = goods.shop_id").
		Where(&Shop{ID: shopId})
	query := tx.ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Find(&goods) })
	if res := tx.Find(&goods); res.Error != nil {
		err = wrapError(res, query)
	}
	fmt.Printf("search shop goods, shop_id: %d\n", shopId)

	return goods, err
}
