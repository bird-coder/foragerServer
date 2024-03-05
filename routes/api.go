/*
 * @Description:
 * @Author: yuanshisan
 * @Date: 2021-12-01 15:47:01
 * @LastEditTime: 2023-02-17 15:35:26
 * @LastEditors: yuanshisan
 */
package routes

import (
	"foragerServer/controller"
	"foragerServer/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func HandleApi(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))

	api := r.Group("/api")
	api.Use(sessions.Sessions("GOSESSID", store))
	// if gin.Mode() == string(constants.PRO) {
	api.Use(middleware.ErrorHandler())
	// }

	handleAuth(api)
	handleUser(api)
	handleProduct(api)
}

func handleAuth(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	auth.POST("/fastLogin", controller.FastLogin)
	auth.POST("/passLogin", controller.PassLogin)
	auth.POST("/sendCode", controller.SendCode)
	auth.POST("/logout", controller.Logout)
}

func handleUser(api *gin.RouterGroup) {
	user := api.Group("/user")
	user.Use(middleware.AuthRequired())
	{
		user.POST("uploadAvatar", controller.UploadAvatar)
		user.GET("/address/list", controller.GetAddressList)
		user.GET("/address/info", controller.GetAddressInfo)
		user.POST("/address/add", controller.CreateAddress)
		user.POST("/address/update", controller.UpdateAddress)
		user.POST("/address/delete", controller.DeleteAddress)
	}
}

func handleProduct(api *gin.RouterGroup) {
	product := api.Group("/product")
	product.GET("/list", controller.GetShopList)
	product.GET("/goods", controller.GetShopGoods)
}
