package controllers

import (
	"net/http"
	"strconv"

	"github.com/Sanket-Ugale/FinSync/internal/models"
	"github.com/gin-gonic/gin"
)

func CreatePortfolio(c *gin.Context) {
	userID := c.GetUint("userID")
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolio := models.Portfolio{
		UserID: userID,
		Name:   input.Name,
	}

	if err := models.CreatePortfolio(&portfolio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create portfolio"})
		return
	}

	c.JSON(http.StatusCreated, portfolio)
}

func GetPortfolios(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolios, err := models.GetPortfoliosByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get portfolios"})
		return
	}

	c.JSON(http.StatusOK, portfolios)
}

func GetPortfolio(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	portfolio, err := models.GetPortfolioByID(uint(portfolioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	if portfolio.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, portfolio)
}

func UpdatePortfolio(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var input struct {
		Name string `json:"name" binding:"required"`
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

	if err := models.UpdatePortfolio(uint(portfolioID), input.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update portfolio"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Portfolio updated successfully"})
}

func DeletePortfolio(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	portfolio, err := models.GetPortfolioByID(uint(portfolioID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	if portfolio.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := models.DeletePortfolio(uint(portfolioID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete portfolio"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Portfolio deleted successfully"})
}
