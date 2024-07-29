package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Sanket-Ugale/FinSync/internal/models"
	"github.com/gin-gonic/gin"
)

func GetPortfolioValue(c *gin.Context) {
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

	value, err := models.CalculatePortfolioValue(uint(portfolioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate portfolio value"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": value})
}

func GetPortfolioReturn(c *gin.Context) {
	userID := c.GetUint("userID")
	portfolioID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	startDate, err := time.Parse("2006-01-02", c.Query("start_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
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

	returnValue, err := models.CalculatePortfolioReturn(uint(portfolioID), startDate, endDate)
	if err != nil {
		log.Printf("Failed to calculate portfolio return: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate portfolio return", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"return": returnValue})
}
