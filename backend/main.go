package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// 初始化数据库
	if err := initDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 初始化Redis
	if err := initRedis(); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	// 创建Gin引擎
	router := gin.Default()

	// 注册中间件
	router.Use(corsMiddleware())
	router.Use(loggerMiddleware())

	// 注册路由
	registerRoutes(router)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server starting on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDatabase() error {
	// TODO: 实现数据库初始化
	fmt.Println("✓ Database initialized")
	return nil
}

func initRedis() error {
	// TODO: 实现Redis初始化
	fmt.Println("✓ Redis initialized")
	return nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("[%s] %s %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP())
		c.Next()
	}
}

func registerRoutes(router *gin.Engine) {
	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "HOHO Miniapp Backend is running",
		})
	})

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 用户相关路由 (待实现)
		users := v1.Group("/users")
		{
			users.POST("/register", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Register endpoint - TODO"})
			})
			users.POST("/login", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Login endpoint - TODO"})
			})
			users.GET("/profile", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Profile endpoint - TODO"})
			})
		}

		// 藏品相关路由 (待实现)
		assets := v1.Group("/assets")
		{
			assets.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "List assets - TODO"})
			})
			assets.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Create asset - TODO"})
			})
			assets.GET("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get asset - TODO"})
			})
		}

		// 交易相关路由 (待实现)
		trades := v1.Group("/trades")
		{
			trades.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "List trades - TODO"})
			})
			trades.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Create trade - TODO"})
			})
		}

		// 积分相关路由 (待实现)
		points := v1.Group("/points")
		{
			points.GET("/balance", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get balance - TODO"})
			})
			points.GET("/history", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Get history - TODO"})
			})
		}

		// 社区事件公示路由 (待实现)
		events := v1.Group("/events")
		{
			events.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "List events - TODO"})
			})
		}
	}

	// 后台管理路由 (待实现)
	admin := router.Group("/admin")
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Admin dashboard - TODO"})
		})
	}
}
