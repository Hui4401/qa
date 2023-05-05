package model

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	UserID     uint   `gorm:"not null;"`           // 回答所属用户Id
	QuestionID uint   `gorm:"not null;"`           // 回答所属问题Id
	Content    string `gorm:"type:text;not null;"` // 内容
	LikeCount  uint   `gorm:"not null;"`           // 点赞数
}
