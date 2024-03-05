package controller

import (
	"fmt"
	"foragerServer/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadAvatar(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		fmt.Print(file.Filename)
		c.SaveUploadedFile(file, "")
	}
}

func GetAddressList(c *gin.Context) {
	user := c.MustGet("user").(*UserSession)
	data := model.GetAddressList(user.UID)
	successResponse(c, data)
}

func GetAddressInfo(c *gin.Context) {
	user := c.MustGet("user").(*UserSession)
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		fmt.Printf("param error, id: %d, err: %s", id, err)
		apiResponse(c, "参数错误", PARAM_ERROR)
		return
	}
	data := model.GetAddressInfo(id, user.UID)
	if data.IsEmpty() {
		apiResponse(c, "地址不存在", PARAM_ERROR)
		return
	}
	successResponse(c, data)
}

func CreateAddress(c *gin.Context) {
	user := c.MustGet("user").(*UserSession)
	var json model.UserAddress
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Printf("param error, params: %v, err: %s", json, err)
		apiResponse(c, "参数错误", PARAM_ERROR)
		return
	}
	json.UID = user.UID
	successResponse(c, nil)
}

func UpdateAddress(c *gin.Context) {
	user := c.MustGet("user").(*UserSession)
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		fmt.Printf("param error, id: %d, err: %s", id, err)
		apiResponse(c, "参数错误", PARAM_ERROR)
		return
	}
	data := model.GetAddressInfo(id, user.UID)
	if data.IsEmpty() {
		apiResponse(c, "地址不存在", PARAM_ERROR)
		return
	}
	var json model.UserAddress
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Printf("param error, params: %v, err: %s", json, err)
		apiResponse(c, "参数错误", PARAM_ERROR)
		return
	}
	successResponse(c, nil)
}

func DeleteAddress(c *gin.Context) {
	user := c.MustGet("user").(*UserSession)
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		fmt.Printf("param error, id: %d, err: %s", id, err)
		apiResponse(c, "参数错误", PARAM_ERROR)
		return
	}
	data := model.GetAddressInfo(id, user.UID)
	if data.IsEmpty() {
		apiResponse(c, "地址不存在", PARAM_ERROR)
		return
	}
	if ok := model.DelAddress(id); !ok {
		apiResponse(c, "删除地址失败", FAIL)
		return
	}
	successResponse(c, nil)
}
