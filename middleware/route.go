/*
 * @Description:
 * @Author: yuanshisan
 * @Date: 2021-12-31 12:47:47
 * @LastEditTime: 2023-02-17 15:43:53
 * @LastEditors: yuanshisan
 */
package middleware

import (
	"encoding/json"
	"fmt"
	"foragerServer/controller"
	zlog "foragerServer/service/logger"
	"foragerServer/util"
	"net/http"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code": controller.LOGIN_EXPIRE,
					"msg":  "登陆已失效",
					"data": nil,
				})
			}
		}()
		session := sessions.Default(c)
		data := session.Get("user")
		user := &controller.UserSession{}
		fmt.Println(data)
		if data == nil {
			panic("session已失效")
		}
		if _, ok := data.(string); !ok {
			panic("session格式错误")
		}
		if err := json.Unmarshal([]byte(data.(string)), user); err != nil {
			panic("session格式错误")
		}
		if reflect.DeepEqual(user, &controller.UserSession{}) {
			panic("用户session信息为空")
		}
		if user.Ip != c.ClientIP() {
			msg := fmt.Sprintf("用户登陆ip异常, last ip: %s, current ip: %s", user.Ip, c.ClientIP())
			panic(msg)
		}
		if user.Expire < util.GetTimestamp() {
			panic("用户登陆token已过期")
		}
		c.Set("user", user)

		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				zlog.Error("server error: [%+v]", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": controller.SERVER_ERROR,
					"msg":  "服务器错误",
					"data": nil,
				})
			}
		}()
		c.Next()
	}
}
