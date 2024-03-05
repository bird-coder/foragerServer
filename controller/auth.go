package controller

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"foragerServer/model"
	"foragerServer/util"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ErrorCode uint16

const (
	SUCCESS        ErrorCode = 0
	FAIL           ErrorCode = 1
	PARAM_ERROR    ErrorCode = 2
	LOGIN_EXPIRE   ErrorCode = 3
	SIGN_ERROR     ErrorCode = 4
	CODE_TIMEOUT   ErrorCode = 5
	INVALID_PARAMS ErrorCode = 10
	SERVER_ERROR   ErrorCode = 500
)

const (
	TEST_PHONE  string = "15026888582"
	CODE_EXPIRE int    = 300
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserSession struct {
	UID     int    `json:"uid"`
	Account string `json:"account"`
	Expire  int    `json:"expire"`
	Ip      string `json:"ip"`
}

//快捷登录
func FastLogin(c *gin.Context) {
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	utype := c.DefaultPostForm("utype", "tel")

	if phone == "" || code == "" {
		apiResponse(c, "参数无效", PARAM_ERROR)
		return
	}
	if len(phone) != 11 || !util.ValidateMobile(phone) {
		apiResponse(c, "手机号错误", FAIL)
		return
	}
	account := fmt.Sprintf("%s@%s", phone, utype)
	user := model.GetUserByAccount(account)
	ip := c.ClientIP()
	timestamp := util.GetTimestamp()
	if user.IsEmpty() {
		//注册
		newUser := &model.User{
			Account:    account,
			Phone:      phone,
			RegIP:      ip,
			RegDate:    timestamp,
			LoginIP:    ip,
			LoginDate:  timestamp,
			LoginTimes: 1,
		}
		uid := model.CreateUser(newUser)
		if uid <= 0 {
			apiResponse(c, "注册失败!", FAIL)
			return
		}
	} else {
		//登陆
		updateLoginData(c, &user)
	}
	if !saveUserSession(c, &user) {
		apiResponse(c, "登陆失败", FAIL)
		return
	}
	apiResponse(c, "", SUCCESS)
}

//密码登陆
func PassLogin(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	utype := c.DefaultPostForm("utype", "tel")

	if phone == "" || password == "" {
		apiResponse(c, "参数无效", PARAM_ERROR)
		return
	}
	if len(phone) != 11 || !util.ValidateMobile(phone) {
		apiResponse(c, "手机号错误", FAIL)
		return
	}
	account := fmt.Sprintf("%s@%s", phone, utype)
	user := model.GetUserByAccount(account)
	if user.IsEmpty() {
		apiResponse(c, "账号未注册", FAIL)
		return
	}
	if user.Password != encryptPassword(password, user.Salt) {
		apiResponse(c, "密码错误", FAIL)
		return
	}
	go updateLoginData(c, &user)

	if !saveUserSession(c, &user) {
		apiResponse(c, "登陆失败", FAIL)
		return
	}
	apiResponse(c, "", SUCCESS)
}

func SendCode(c *gin.Context) {
	phone := c.PostForm("phone")
	if len(phone) != 11 || !util.ValidateMobile(phone) {
		apiResponse(c, "手机号错误", FAIL)
		return
	}
	random, _ := rand.Int(rand.Reader, big.NewInt(899999))
	code := random.Int64() + 100000
	if strings.Contains(TEST_PHONE, phone) {
		code = 123456
	}
	sendCode := &model.SendCode{
		Phone: phone,
		Code:  int(code),
		Ltime: util.GetTimestamp(),
	}
	if !model.InsertSendCode(sendCode) {
		apiResponse(c, "验证码发送失败，请重试!", FAIL)
		return
	}
	apiResponse(c, "验证码已发送请查收!", SUCCESS)
}

func Logout(c *gin.Context) {
	if !removeUserSession(c) {
		apiResponse(c, "操作失败", FAIL)
		return
	}
	apiResponse(c, "", SUCCESS)
}

//user信息写入session
func saveUserSession(c *gin.Context, user *model.User) bool {
	userSession := &UserSession{
		UID:     user.UID,
		Account: user.Account,
		Expire:  util.GetExpireTime(7),
		Ip:      c.ClientIP(),
	}
	data, _ := json.Marshal(userSession)
	session := sessions.Default(c)
	session.Set("user", string(data))
	if err := session.Save(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//user信息从session移除
func removeUserSession(c *gin.Context) bool {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		session.Delete("user")
		if err := session.Save(); err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}

func updateLoginData(c *gin.Context, user *model.User) bool {
	if user.IsEmpty() {
		return false
	}
	user.LoginTimes = user.LoginTimes + 1
	user.LoginDate = util.GetTimestamp()
	user.LoginIP = c.ClientIP()
	model.UpdateUser(user)
	return true
}

//验证码校验
func checkPhoneCode(phone string, code int) bool {
	sendCode := model.GetSendCode(phone)
	if sendCode.Code != code {
		return false
	}
	if sendCode.Ltime+300 < util.GetTimestamp() {
		return false
	}
	return true
}

func successResponse(c *gin.Context, data interface{}) {
	res := &Response{
		Code: int(SUCCESS),
		Msg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func apiResponse(c *gin.Context, msg string, code ErrorCode) {
	res := &Response{
		Code: int(code),
		Msg:  msg,
	}
	c.JSON(http.StatusOK, res)
}

func createSalt(n int) string {
	letters := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	salt := make([]byte, n)
	for i := range salt {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		salt[i] = letters[idx.Int64()]
	}
	return string(salt)
}

func encryptPassword(password string, salt string) string {
	m1 := md5.New()
	m1.Write([]byte(password))
	passMd5 := hex.EncodeToString(m1.Sum(nil))

	m2 := md5.New()
	m2.Write([]byte(passMd5))
	m2.Write([]byte(salt))
	return hex.EncodeToString(m2.Sum(nil))
}
