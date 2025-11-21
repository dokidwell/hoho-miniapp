package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hoho-miniapp/backend/config"
	"hoho-miniapp/backend/database"
	"hoho-miniapp/backend/handlers"
	"hoho-miniapp/backend/middleware"
	"hoho-miniapp/backend/services"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// 初始化配置
	config.InitConfig()
	fmt.Println("✅ Configuration initialized")

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
	return database.InitDatabase()
}

func initRedis() error {
	return database.InitRedis()
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
	// 初始化服务和处理器
	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(userService)
	assetService := services.NewAssetService()
	assetHandler := handlers.NewAssetHandler(assetService)
	eventService := services.NewEventService()
	eventHandler := handlers.NewEventHandler(eventService)
	jingtanService := services.NewJingtanService()
	jingtanHandler := handlers.NewJingtanHandler(jingtanService)
	tradeService := services.NewTradeService()
	tradeHandler := handlers.NewTradeHandler(tradeService)
	uploadHandler := handlers.NewUploadHandler()
	airdropService := services.NewAirdropService()
	
	// 初始化新增服务和处理器
	creationService := services.NewCreationService()
	creationHandler := handlers.NewCreationHandler(creationService)
	taskService := services.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)
	announcementService := services.NewAnnouncementService()
	announcementHandler := handlers.NewAnnouncementHandler(announcementService)
	offerService := services.NewOfferService()
	offerHandler := handlers.NewOfferHandler(offerService)
	platformAccountService := services.NewPlatformAccountService()
	platformAccountHandler := handlers.NewPlatformAccountHandler(platformAccountService)
	
	// 初始化管理员服务和处理器
	adminService := services.NewAdminService()
	adminHandler := handlers.NewAdminHandler(adminService, assetService, airdropService)
	adminCreationHandler := handlers.NewAdminCreationHandler(creationService)
	adminTaskHandler := handlers.NewAdminTaskHandler(taskService)
	adminAnnouncementHandler := handlers.NewAdminAnnouncementHandler(announcementService)
	adminConfigHandler := handlers.NewAdminConfigHandler()

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "HOHO Miniapp Backend is running",
		})
	})

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 用户相关路由 (无需认证)
		users := v1.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
		}

		// 需要认证的路由
		auth := v1.Group("/")
		auth.Use(middleware.AuthMiddleware())
		{
			// 用户相关路由
			auth.GET("/users/profile", userHandler.GetProfile)
			auth.PUT("/users/profile", userHandler.UpdateProfile)
			auth.POST("/users/verify-identity", userHandler.VerifyIdentity)
			auth.GET("/users/points", userHandler.GetPoints)

			// 藏品相关路由
			assets := auth.Group("/assets")
			{
				assets.POST("", assetHandler.SubmitMintRequest) // 提交铸造请求
			}

			// 交易相关路由
			trades := auth.Group("/trades")
			{
				trades.POST("/execute", tradeHandler.ExecuteTrade)
				trades.GET("/history", tradeHandler.GetTradeHistory)
				trades.GET("/:id", tradeHandler.GetTradeDetail)
			}

			// 挂售相关路由
			listings := auth.Group("/listings")
			{
				listings.POST("", tradeHandler.CreateListing)
				listings.DELETE("/:id", tradeHandler.CancelListing)
			}

			// 我的相关
			my := auth.Group("/my")
			{
				my.GET("/listings", tradeHandler.GetMyListings)
				my.GET("/assets", assetHandler.GetMyAssets)
			}

			// 上传相关
			upload := auth.Group("/upload")
			{
				upload.POST("", uploadHandler.UploadFile)
				upload.GET("/cos-credentials", uploadHandler.GetCOSCredentials)
				upload.GET("/oss-credentials", uploadHandler.GetOSSCredentials)
			}

			// 积分相关路由 (待实现)
			_ = auth.Group("/points")
			// {
			// 	points.GET("/balance", pointHandler.GetBalance)
			// 	points.GET("/history", pointHandler.GetHistory)
			// }

			// 社区事件路由
			events := auth.Group("/events")
			{
				events.GET("", eventHandler.ListEvents)
				events.GET("/:id", eventHandler.GetEventDetail)
			}

				// 鲸探API路由
				jingtan := auth.Group("/jingtan")
				{
					jingtan.POST("/bind", jingtanHandler.BindAccount)
					jingtan.DELETE("/unbind", jingtanHandler.UnbindAccount)
					jingtan.POST("/sync", jingtanHandler.SyncAssets)
					jingtan.GET("/assets", jingtanHandler.GetAssets)
				}
				
				// 创作相关路由
				creations := auth.Group("/creations")
				{
					creations.POST("", creationHandler.SubmitCreation)
					creations.GET("", creationHandler.GetMyCreations)
					creations.GET("/:id", creationHandler.GetCreationDetail)
				}
				
				// 任务相关路由
				tasks := auth.Group("/tasks")
				{
					tasks.GET("", taskHandler.GetTaskList)
					tasks.POST("/:id/complete", taskHandler.CompleteTask)
					tasks.POST("/:id/claim", taskHandler.ClaimReward)
				}
				
				// 出价/心愿单相关路由
				offers := auth.Group("/offers")
				{
					offers.POST("", offerHandler.CreateOffer)
					offers.GET("", offerHandler.GetMyOffers)
					offers.DELETE("/:id", offerHandler.CancelOffer)
					offers.POST("/:id/accept", offerHandler.AcceptOffer)
				}
			}

		// 公开的藏品路由
		assetsPublic := v1.Group("/assets")
		{
			assetsPublic.GET("", assetHandler.ListAssets)
			assetsPublic.GET("/:id", assetHandler.GetAssetDetail)
		}

		// 公开的集换路由
		listingsPublic := v1.Group("/listings")
		{
			listingsPublic.GET("", tradeHandler.ListListings)
			listingsPublic.GET("/:id", tradeHandler.GetListingDetail)
		}

			// 公开的事件路由
			eventsPublic := v1.Group("/events")
			{
				eventsPublic.GET("", eventHandler.ListEvents)
				eventsPublic.GET("/:id", eventHandler.GetEventDetail)
			}
			
			// 公开的公告路由
			announcementsPublic := v1.Group("/announcements")
			{
				announcementsPublic.GET("", announcementHandler.GetAnnouncementList)
				announcementsPublic.GET("/:id", announcementHandler.GetAnnouncementDetail)
			}
			
			// 公开的平台账户路由
			platformAccountPublic := v1.Group("/platform-account")
			{
				platformAccountPublic.GET("", platformAccountHandler.GetAccountInfo)
				platformAccountPublic.GET("/transactions", platformAccountHandler.GetTransactions)
			}
		}

	// 注册自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
	})

	// 加载HTML模板
	router.LoadHTMLGlob("templates/*.html")

	// 后台管理路由
	admin := router.Group("/admin")
	{
		// 登录页面
		admin.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin_login.html", gin.H{})
		})
		// 登录API
		admin.POST("/login", adminHandler.Login)

		// 需要认证的路由
		authAdmin := admin.Group("/")
		authAdmin.Use(middleware.AdminAuthMiddleware())
		{
			authAdmin.GET("/profile", adminHandler.GetProfile)
			authAdmin.GET("/dashboard", func(c *gin.Context) {
				// 模拟数据
				data := gin.H{
					"Title":            "仪表盘",
					"ActiveMenu":       "dashboard",
					"TotalUsers":       1234,
					"TotalAssets":      567,
					"PendingReviews":   12,
					"TodayTradeVolume": "12,345.67890123",
				}
				c.HTML(http.StatusOK, "admin_dashboard.html", data)
			})

			// 用户管理路由
			users := authAdmin.Group("/users")
			{
				users.GET("", adminHandler.ListUsersPage)
				// users.GET("/:id", adminHandler.GetUserDetailPage) // 用户详情页
				users.PUT("/:id/status", adminHandler.UpdateUserStatus) // 禁用/解禁 API
			}

			// 藏品审核路由
			assetsReview := authAdmin.Group("/review/assets")
			{
				assetsReview.GET("", adminHandler.ListAssetReviewPage)
				assetsReview.PUT("/:id", adminHandler.ReviewAsset)
			}

				// 空投管理路由
				airdrop := authAdmin.Group("/airdrop")
				{
					airdrop.POST("/points", adminHandler.AirdropPoints)
					airdrop.POST("/asset", adminHandler.AirdropAsset)
				}
				
				// 创作审核管理路由
				creationsAdmin := authAdmin.Group("/creations")
				{
					creationsAdmin.GET("", adminCreationHandler.GetCreationList)
					creationsAdmin.GET("/:id", adminCreationHandler.GetCreationDetail)
					creationsAdmin.POST("/:id/approve", adminCreationHandler.ApproveCreation)
					creationsAdmin.POST("/:id/reject", adminCreationHandler.RejectCreation)
				}
				
				// 任务管理路由
				tasksAdmin := authAdmin.Group("/tasks")
				{
					tasksAdmin.GET("", adminTaskHandler.GetTaskList)
					tasksAdmin.POST("", adminTaskHandler.CreateTask)
					tasksAdmin.PUT("/:id", adminTaskHandler.UpdateTask)
					tasksAdmin.DELETE("/:id", adminTaskHandler.DeleteTask)
					tasksAdmin.POST("/:id/toggle", adminTaskHandler.ToggleTask)
				}
				
				// 公告管理路由
				announcementsAdmin := authAdmin.Group("/announcements")
				{
					announcementsAdmin.GET("", adminAnnouncementHandler.GetAnnouncementList)
					announcementsAdmin.POST("", adminAnnouncementHandler.CreateAnnouncement)
					announcementsAdmin.PUT("/:id", adminAnnouncementHandler.UpdateAnnouncement)
					announcementsAdmin.DELETE("/:id", adminAnnouncementHandler.DeleteAnnouncement)
					announcementsAdmin.POST("/:id/toggle-pin", adminAnnouncementHandler.TogglePin)
				}
				
				// 系统配置管理路由
				configAdmin := authAdmin.Group("/config")
				{
					configAdmin.GET("", adminConfigHandler.GetConfig)
					configAdmin.PUT("", adminConfigHandler.UpdateConfig)
				}
			}
		}
		
		// 静态文件服务（管理后台）
		router.Static("/admin-ui", "./admin")
	}
