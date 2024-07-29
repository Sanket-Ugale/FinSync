package main

import (
	"log"

	"github.com/Sanket-Ugale/FinSync/internal/controllers"
	"github.com/Sanket-Ugale/FinSync/internal/middleware"
	"github.com/Sanket-Ugale/FinSync/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	models.InitDB()

	// Initialize Redis
	models.InitRedis()

	// Set up Gin router
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())
	// r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	// website routes
	r.Static("/templates", "./templates")
	// website routes
	r.Static("/assets", "./assets")

	// Serve index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./templates/index.html")
	})

	// Routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/verify-otp", controllers.VerifyOTP)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// User routes
		api.GET("/user", controllers.GetUserDetails)
		api.PUT("/user", controllers.UpdateUser)

		// Portfolio routes
		api.POST("/portfolio", controllers.CreatePortfolio)
		api.GET("/portfolio", controllers.GetPortfolios)
		api.GET("/portfolio/:id", controllers.GetPortfolio)
		api.PUT("/portfolio/:id", controllers.UpdatePortfolio)
		api.DELETE("/portfolio/:id", controllers.DeletePortfolio)

		// Asset routes
		api.POST("/portfolio/:id/asset", controllers.AddAsset)
		api.PUT("/portfolio/:id/asset/:assetId", controllers.UpdateAsset)
		api.DELETE("/portfolio/:id/asset/:assetId", controllers.DeleteAsset)
		api.GET("/portfolio/:id/asset/:assetId", controllers.GetAsset)

		// Analytics routes
		api.GET("/portfolio/:id/value", controllers.GetPortfolioValue)
		api.GET("/portfolio/:id/return", controllers.GetPortfolioReturn)
	}

	r.Run("0.0.0.0:80")
}
