package main

import (
	"log"
	"os"

	"github.com/Sanket-Ugale/FinSync/internal/controllers"
	"github.com/Sanket-Ugale/FinSync/internal/middleware"
	"github.com/Sanket-Ugale/FinSync/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (optional in production)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	models.InitDB()

	// Initialize Redis (optional for some deployments)
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

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"service": "FinSync API",
			"version": "1.0.0",
		})
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

	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)
}
