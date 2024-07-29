package controllers

import (
	"net/http"
	"strconv"

	"github.com/Sanket-Ugale/FinSync/internal/models"
	"github.com/gin-gonic/gin"
)

func AddAsset(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var input struct {
		Name     string  `json:"name" binding:"required"`
		Type     string  `json:"type" binding:"required"`
		Quantity float64 `json:"quantity" binding:"required"`
		Value    float64 `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio, err := models.GetPortfolioByID(uint(portfolioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	if portfolio.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	asset := models.Asset{
		PortfolioID: uint(portfolioID),
		Name:        input.Name,
		Type:        input.Type,
		Quantity:    input.Quantity,
		Value:       input.Value,
	}

	if err := models.CreateAsset(&asset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add asset"})
		return
	}

	c.JSON(http.StatusCreated, asset)
}

func UpdateAsset(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	assetID, _ := strconv.ParseUint(c.Param("assetId"), 10, 64)

	var input struct {
		Name     string  `json:"name"`
		Type     string  `json:"type"`
		Quantity float64 `json:"quantity"`
		Value    float64 `json:"value"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio, err := models.GetPortfolioByID(uint(portfolioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	if portfolio.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	asset, err := models.GetAssetByID(uint(assetID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if asset.PortfolioID != uint(portfolioID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := models.UpdateAsset(uint(assetID), input.Name, input.Type, input.Quantity, input.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset updated successfully"})
}

// controllers/asset.go (continued)

func DeleteAsset(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	assetID, _ := strconv.ParseUint(c.Param("assetId"), 10, 64)

	portfolio, err := models.GetPortfolioByID(uint(portfolioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	if portfolio.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	asset, err := models.GetAssetByID(uint(assetID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if asset.PortfolioID != uint(portfolioID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := models.DeleteAsset(uint(assetID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset deleted successfully"})
}

func GetAsset(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	assetID, _ := strconv.ParseUint(c.Param("assetId"), 10, 64)

	portfolio, err := models.GetPortfolioByID(uint(portfolioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	if portfolio.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	asset, err := models.GetAssetByID(uint(assetID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if asset.PortfolioID != uint(portfolioID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, asset)
}
