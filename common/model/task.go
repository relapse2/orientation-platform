package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserId          string `gorm:"index"`
	Title           string
	Type            int32 `gorm:"index"`
	Text            string
	ImageUrl        string
	AnswerUrl       string
	Bonus           int64
	State           int64  `gorm:"index"` //0为未完成或失败，1为转交人工审核，2为处理成功
	ReferenceAnswer string //参考图片答案
}
