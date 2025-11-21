// routes_v2.go - 新版路由配置（统一术语为"作品"）
package main

import (
	"hoho-api/handlers"
	"hoho-api/middleware"

	"github.com/gin-gonic/gin"
)

func setupRoutesV2(r *gin.Engine) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 公开接口
	public := r.Group("/api/v1")
	{
		// 用户相关
		public.POST("/register", handlers.Register)
		public.POST("/login", handlers.UserLogin)
		
		// 公告
		public.GET("/announcements", handlers.GetAnnouncements)
		public.GET("/announcements/:id", handlers.GetAnnouncementDetail)
		
		// 作品（公开）
		public.GET("/artworks", handlers.GetArtworks) // 需要创建
		public.GET("/artworks/:id", handlers.GetArtworkDetail) // 需要创建
		
		// 平台账户（阳光账户）- 公开透明
		public.GET("/platform/account", handlers.GetPlatformAccount)
		public.GET("/platform/transactions", handlers.GetPlatformTransactions)
	}

	// 需要认证的接口
	auth := r.Group("/api/v1")
	auth.Use(middleware.AuthMiddleware())
	{
		// 用户信息
		auth.GET("/user/profile", handlers.GetUserProfile)
		auth.PUT("/user/profile", handlers.UpdateUserProfile)
		
		// 积分
		auth.GET("/points/balance", handlers.GetPointsBalance)
		auth.GET("/points/transactions", handlers.GetPointsTransactions)
		
		// 创作
		auth.POST("/creations", handlers.CreateCreation)
		auth.GET("/creations", handlers.GetMyCreations)
		auth.GET("/creations/:id", handlers.GetCreationDetail)
		
		// 任务
		auth.GET("/tasks", handlers.GetTasks)
		auth.POST("/tasks/signin", handlers.DailySignIn)
		auth.POST("/tasks/claim/:completion_id", handlers.ClaimTaskReward)
		
		// 出价（心愿单）
		auth.POST("/offers", handlers.CreateOffer)
		auth.GET("/offers/sent", handlers.GetMyOffers)
		auth.GET("/offers/received", handlers.GetReceivedOffers)
		auth.POST("/offers/:id/accept", handlers.AcceptOffer)
		auth.POST("/offers/:id/reject", handlers.RejectOffer)
		auth.DELETE("/offers/:id", handlers.CancelOffer)
		
		// 通知
		auth.GET("/notifications", handlers.GetNotifications) // 需要创建
		auth.PUT("/notifications/:id/read", handlers.MarkNotificationRead) // 需要创建
		
		// 我的作品
		auth.GET("/my/artworks", handlers.GetMyArtworks) // 需要创建
		
		// 交易
		auth.GET("/trades", handlers.GetMyTrades) // 需要创建
	}

	// 管理员接口
	admin := r.Group("/admin")
	admin.Use(middleware.AdminAuthMiddleware())
	{
		// 创作审核
		admin.GET("/creations/pending", handlers.GetPendingCreations) // 需要创建
		admin.POST("/creations/:id/approve", handlers.ApproveCreation) // 需要创建
		admin.POST("/creations/:id/reject", handlers.RejectCreation) // 需要创建
		admin.POST("/creations/:id/publish", handlers.PublishCreation) // 需要创建
		
		// 任务管理
		admin.GET("/tasks", handlers.GetAllTasksAdmin) // 需要创建
		admin.POST("/tasks", handlers.CreateTask) // 需要创建
		admin.PUT("/tasks/:id", handlers.UpdateTask) // 需要创建
		admin.DELETE("/tasks/:id", handlers.DeleteTask) // 需要创建
		
		// 公告管理
		admin.POST("/announcements", handlers.CreateAnnouncement) // 需要创建
		admin.PUT("/announcements/:id", handlers.UpdateAnnouncement) // 需要创建
		admin.POST("/announcements/:id/publish", handlers.PublishAnnouncement) // 需要创建
		admin.DELETE("/announcements/:id", handlers.DeleteAnnouncement) // 需要创建
		
		// 系统配置
		admin.GET("/configs", handlers.GetSystemConfigs) // 需要创建
		admin.PUT("/configs/:key", handlers.UpdateSystemConfig) // 需要创建
	}
}
