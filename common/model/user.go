package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique;index"`
	Password     string `gorm:"not null"`
	IdCard       string `gorm:"not null;index"`
	Sex          int64
	Collage      string `gorm:"index"`
	Point        int64
	RegAddressID uint
	RegAddress   Location `gorm:"foreignKey:RegAddressID;references:ID"`
	AvatarUrl    string   `gorm:"not null"`
	CharacterUrl string
}

type Location struct {
	gorm.Model
	Address   string //地域名
	Latitude  float64
	Longitude float64
}

type Admin struct {
	gorm.Model
	Username string `gorm:"not null;unique;index"`
	Password string `gorm:"not null"`
}

type Student struct {
	gorm.Model
	Name       string
	Collage    string //可能后续做一下不同院校之间分组？
	IdCard     string `gorm:"not null;index"`
	IfVerified bool   //0为未被认证，1为已被认证
}

const (
	PopularUserStandard = 1000 // 拥有超过 1000 积分的用户成为热门用户，有特殊处理
)

func IsPopularUser(point int64) bool {
	return point >= PopularUserStandard
}
