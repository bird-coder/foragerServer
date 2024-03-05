package util

import (
	"fmt"
	"math"
	"regexp"
	"time"
)

//获取当前时间戳
func GetTimestamp() int {
	return int(time.Now().Unix())
}

//获取几天后/前时间戳
func GetExpireTime(day int) int {
	return int(time.Now().Add(time.Duration(day * 24 * int(time.Hour))).Unix())
}

//获取当前日期
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

//获取当前日期时间
func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//时间戳转日期
func ParseToDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//日期转时间戳
func ParseToTimestamp(date string) int64 {
	tm, _ := time.Parse("2006-01-02 15:04:05", date)
	return tm.Unix()
}

func ValidateMobile(phone string) bool {
	if _, err := regexp.Match(`^0?(1[3-9][0-9])[0-9]{8}$`, []byte(phone)); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ValidateIp(ip string) bool {
	if _, err := regexp.Match(`^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, []byte(ip)); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetLocationRange(lat float64, lng float64, limit float64) (float64, float64, float64, float64) {
	distancePerDegree := 111.1 * 1000
	diffLat := math.Round(limit/(math.Cos(lat)*distancePerDegree)*1000000) / 1000000 / 2
	diffLng := math.Round(limit/distancePerDegree*1000000) / 1000000 / 2
	minLat := lat - diffLat
	maxLat := lat + diffLat
	minLng := lng - diffLng
	maxLng := lng + diffLng
	return minLat, maxLat, minLng, maxLng
}

func CalcDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	radLat1 := lat1 * math.Pi / 180.0
	radLat2 := lat2 * math.Pi / 180.0
	x := radLat1 - radLat2
	y := (lng1 - lng2) * math.Pi / 180.0
	s := math.Asin(math.Sqrt(math.Pow(math.Sin(x/2), 2))) + math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(y/2), 2)
	return s * 6378137
}
