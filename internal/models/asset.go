package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	PortfolioID uint
	Name        string
	Type        string
	Quantity    float64
	Value       float64
}

type AssetHistory struct {
	gorm.Model
	AssetID uint
	Value   float64
	Date    time.Time
}

func CreateAsset(asset *Asset) error {
	return DB.Create(asset).Error
}

func GetAssetByID(id uint) (*Asset, error) {
	var asset Asset
	err := DB.First(&asset, id).Error
	return &asset, err
}

func UpdateAsset(id uint, name string, assetType string, quantity float64, value float64) error {
	return DB.Model(&Asset{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":     name,
		"type":     assetType,
		"quantity": quantity,
		"value":    value,
	}).Error
}

func DeleteAsset(id uint) error {
	return DB.Delete(&Asset{}, id).Error
}

func CalculatePortfolioValue(portfolioID uint) (float64, error) {
	var totalValue float64
	err := DB.Model(&Asset{}).Where("portfolio_id = ?", portfolioID).Select("SUM(value * quantity)").Scan(&totalValue).Error
	return totalValue, err
}

func CalculatePortfolioReturn(portfolioID uint, startDate time.Time, endDate time.Time) (float64, error) {
	var startValue, endValue float64

	err := DB.Model(&AssetHistory{}).
		Joins("JOIN assets ON asset_histories.asset_id = assets.id").
		Where("assets.portfolio_id = ? AND asset_histories.date = ?", portfolioID, startDate).
		Select("SUM(asset_histories.value * assets.quantity)").
		Scan(&startValue).Error
	if err != nil {
		return 0, err
	}

	err = DB.Model(&AssetHistory{}).
		Joins("JOIN assets ON asset_histories.asset_id = assets.id").
		Where("assets.portfolio_id = ? AND asset_histories.date = ?", portfolioID, endDate).
		Select("SUM(asset_histories.value * assets.quantity)").
		Scan(&endValue).Error
	if err != nil {
		return 0, err
	}

	return (endValue - startValue) / startValue, nil
}
