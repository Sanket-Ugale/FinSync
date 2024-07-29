package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Name       string
	IsActive   bool `gorm:"default:false"`
	Portfolios []Portfolio
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetUserByID(id uint) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	return &user, err
}

func UpdateUser(id uint, name string) error {
	return DB.Model(&User{}).Where("id = ?", id).Update("name", name).Error
}

func UserExists(email string) bool {
	var count int64
	DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func ActivateUser(email string) error {
	return DB.Model(&User{}).Where("email = ?", email).Update("is_active", true).Error
}
