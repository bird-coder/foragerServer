package controller

import (
	"fmt"
	"foragerServer/model"
	"foragerServer/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type LocationRange struct {
	Lat      float64 `form:"lat"`
	Lng      float64 `form:"lng"`
	Distance float64 `form:"distance"`
}

func GetShopList(c *gin.Context) {
	var form LocationRange
	if err := c.ShouldBind(&form); err != nil {
		fmt.Printf("param error, params: %v, err: %s", form, err)
		apiResponse(c, "参数错误", PARAM_ERROR)
		return
	}
	minLat, maxLat, minLng, maxLng := util.GetLocationRange(form.Lat, form.Lng, form.Distance)
	shopList, err := model.GetShopListByRange(minLat, maxLat, minLng, maxLng)
	if err != nil {
		panic(errors.WithMessage(err, "GetShopList failed"))
	}
	var data []map[string]interface{}
	for _, shop := range shopList {
		tmp := make(map[string]interface{})
		tmp["thumb"] = shop.Thumb
		tmp["title"] = shop.Title
		tmp["score"] = shop.Score
		tmp["lat"] = shop.Lat
		tmp["lng"] = shop.Lng
		tmp["distance"] = util.CalcDistance(form.Lat, form.Lng, shop.Lat, shop.Lng)
		tmp["tags"] = []string{"111", "222"}
		data = append(data, tmp)
	}
	successResponse(c, data)
}

func GetShopGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		fmt.Printf("param error, id: %d, err: %s", id, err)
		apiResponse(c, "参数错误", PARAM_ERROR)
		return
	}
	goods, err := model.GetShopGoods(1)
	if err != nil {
		panic(errors.WithMessage(err, "GetShopGoods failed"))
	}
	if len(goods) == 0 {
		apiResponse(c, "店铺商品为空", PARAM_ERROR)
		return
	}
	successResponse(c, goods)
}
