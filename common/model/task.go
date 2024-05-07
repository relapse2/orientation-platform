package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserId     string `gorm:"index"`
	AnswerUrl  string
	State      int64 `gorm:"index"` //0为未完成或失败，1为转交人工审核，2为处理成功
	TaskInfoID uint
	TaskInfo   SetUpTask
}

type SetUpTask struct {
	gorm.Model
	Collage        string //用于分类
	Title          string
	Text           string
	ImageUrl       string //任务图片
	ReferAnswerUrl string //参考图片答案
	Bonus          int64
	Type           int64 `gorm:"index"` //任务类型
}
