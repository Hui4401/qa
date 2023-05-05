package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/Hui4401/qa/storage/mysql"
)

const (
	// 密码加密级别
	passwordCost = bcrypt.DefaultCost
)

type User struct {
	gorm.Model
	Username    string       // 用户名
	Password    string       // 密码
	UserProfile *UserProfile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联用户信息
	Questions   []Question   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联问题信息
	Answers     []Answer     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联回答信息
}

type UserProfile struct {
	gorm.Model
	UserID      uint
	Nickname    string // 昵称
	Email       string // 邮箱
	Avatar      string // 头像
	Description string // 个人描述
}

type UserDao struct {
	sqlClient *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{
		sqlClient: mysql.GetClient(),
	}
}

func (d *UserDao) CreateUser(u *User) error {
	if err := d.sqlClient.Create(u).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID 没有记录时返回 nil, nil
func (d *UserDao) GetUserByID(ID uint) (*User, error) {
	user := &User{}
	res := d.sqlClient.First(user, ID)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

// GetUserByUsername 没有记录时返回 nil, nil
func (d *UserDao) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	res := d.sqlClient.Where("username = ?", username).First(user)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (d *UserDao) GetUserProfileByID(ID interface{}) (*UserProfile, error) {
	profile := &UserProfile{}
	res := d.sqlClient.Where("user_id = ?", ID).First(profile)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return profile, nil
}

func (d *UserDao) GetUserProfileByUsername(username string) (*UserProfile, error) {
	profile := &UserProfile{}
	res := d.sqlClient.Where("username = ?", username).First(profile)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return profile, nil
}
