package model

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	UserID      uint     `gorm:"not null;"`                                     // 问题所属用户Id
	Title       string   `gorm:"not null;"`                                     // 标题
	Content     string   `gorm:"type:text"`                                     // 内容
	AnswerCount uint     `gorm:"default:0"`                                     // 回答总数
	Answers     []Answer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联回答信息
}
