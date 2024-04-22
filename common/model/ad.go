package model

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	Owner     string `gorm:"not null"`
	Name      string `gorm:"not null"`
	ClickNum  int64
	ClickRate float64
	ImageUrl  string `gorm:"not null"`
	SiteUrl   string `gorm:"not null"`
	TargetNum int64
	Target    Target //最多给三个关键词
}

type Target struct {
	gorm.Model
	District string
	Sex      string
	Collage  string
}
