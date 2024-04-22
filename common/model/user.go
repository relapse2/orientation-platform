package model

import "gorm.io/gorm"

type Location struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string //地域名
	Latitude  float64
	Longitude float64
}

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique;index"`
	Password     string `gorm:"not null"`
	Sex          string
	Collage      string `gorm:"index"`
	Point        int64
	RegAddress   Location `gorm:"not null"`
	AvatarUrl    string   `gorm:"not null"`
	CharacterUrl string
}

type Admin struct {
	gorm.Model
	Username string `gorm:"not null;unique;index"`
	Password string `gorm:"not null"`
}

type Student struct {
	gorm.Model
	name     string
	Collage  string
	IdCard   string `gorm:"not null;index"`
	ImageUrl string
}
