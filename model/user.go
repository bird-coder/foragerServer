package model

import (
	"fmt"
)

func GetUserById(id int) User {
	var user User
	if res := daoModel.Db.Where(&User{UID: id}).First(&user); res.Error != nil {
		fmt.Printf("search user failed, id: %d, err: %s\n", id, res.Error)
	}
	fmt.Println(user)
	return user
}

func GetUserByAccount(account string) User {
	var user User
	if res := daoModel.Db.Where(&User{Account: account}).First(&user); res.Error != nil {
		fmt.Printf("search user failed, account: %s, err: %s\n", account, res.Error)
	}
	fmt.Println(user)
	return user
}

func GetUserByPhone(phone string) User {
	var user User
	if res := daoModel.Db.Where(&User{Phone: phone}).First(&user); res.Error != nil {
		fmt.Printf("search user failed, phone: %s, err: %s\n", phone, res.Error)
	}
	fmt.Println(user)
	return user
}

func CreateUser(user *User) int {
	tx := daoModel.Db.Begin()
	if res := tx.Create(&user); res.Error != nil {
		fmt.Printf("create user failed, err: %s\n", res.Error)
		tx.Rollback()
		return 0
	}
	if res := tx.Create(&UserField{UID: user.UID}); res.Error != nil {
		fmt.Printf("create user failed, err: %s\n", res.Error)
		tx.Rollback()
		return 0
	}
	tx.Commit()
	return user.UID
}

func UpdateUser(user *User) bool {
	if res := daoModel.Db.Save(&user); res.Error != nil {
		fmt.Printf("update user failed, err: %s\n", res.Error)
		return false
	}
	return true
}

func InsertSendCode(sendCode *SendCode) bool {
	if res := daoModel.Db.Create(&sendCode); res.Error != nil {
		fmt.Printf("insert send code failed, err: %s\n", res.Error)
		return false
	}
	return true
}

func GetSendCode(phone string) SendCode {
	var sendCode SendCode
	if res := daoModel.Db.Where(&SendCode{Phone: phone}).First(&sendCode); res.Error != nil {
		fmt.Printf("search send code failed, phone: %s, err: %s\n", phone, res.Error)
	}
	return sendCode
}

func GetAddressList(uid int) []UserAddress {
	var addresses []UserAddress
	if res := daoModel.Db.Where(&UserAddress{UID: uid}).Find(&addresses); res.Error != nil {
		fmt.Printf("search address list failed, err: %s\n", res.Error)
	}
	fmt.Println(addresses)
	return addresses
}

func GetAddressInfo(id int, uid int) UserAddress {
	var address UserAddress
	if res := daoModel.Db.Where(&UserAddress{ID: id, UID: uid}).First(&address); res.Error != nil {
		fmt.Printf("search address info failed, id: %d, err: %s\n", id, res.Error)
	}
	fmt.Println(address)
	return address
}

func AddAddress(address *UserAddress) int {
	if res := daoModel.Db.Create(&address); res.Error != nil {
		fmt.Printf("add address failed, err: %s\n", res.Error)
		return 0
	}
	return address.ID
}

func EditAddress(id int, address map[string]interface{}) bool {
	if res := daoModel.Db.Where(&UserAddress{ID: id}).Updates(address); res.Error != nil {
		fmt.Printf("edit address failed, err: %s\n", res.Error)
		return false
	}
	return true
}

func DelAddress(id int) bool {
	if res := daoModel.Db.Delete(&UserAddress{}, id); res.Error != nil {
		fmt.Printf("delete address failed, err: %s\n", res.Error)
		return false
	}
	return true
}
