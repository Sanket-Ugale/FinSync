package models

import (
	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	UserID uint
	Name   string
	Assets []Asset
}

func CreatePortfolio(portfolio *Portfolio) error {
	return DB.Create(portfolio).Error
}

func GetPortfoliosByUserID(userID uint) ([]Portfolio, error) {
	var portfolios []Portfolio
	err := DB.Where("user_id = ?", userID).Find(&portfolios).Error
	return portfolios, err
}

func GetPortfolioByID(id uint) (*Portfolio, error) {
	var portfolio Portfolio
	err := DB.First(&portfolio, id).Error
	return &portfolio, err
}

func UpdatePortfolio(id uint, name string) error {
	return DB.Model(&Portfolio{}).Where("id = ?", id).Update("name", name).Error
}

func DeletePortfolio(id uint) error {
	return DB.Delete(&Portfolio{}, id).Error
}
