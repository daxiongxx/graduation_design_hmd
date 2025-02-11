package router

import (
	"forum/controller"
	"forum/logger"
	"forum/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	// 设置中间件 recovery 中间件会recovery项目中kennel会出现的panic
	router.Use(logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.RateLimitMiddleware(2*time.Second, 40))
	router.Static("/images", "./images")

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.UserRegisterController)
		user.POST("/login", controller.UserLoginController)
	}
	article := router.Group("/api/article")
	{
		article.POST("/create", controller.ArticleCreateController)
		article.POST("/update", controller.ArticleUpdateController)
		article.GET("/getById/:aid", controller.ArticleGetController)
		article.DELETE("/delete/:aid", controller.ArticleDeleteController)
		article.POST("/list", controller.ArticleGetListController)
		article.POST("/category/list", controller.ArticleCategoryGetListController)
	}
	topic := router.Group("/api/topic")
	{
		topic.POST("/create", controller.TopicCreateController)
		topic.POST("/update", controller.TopicUpdateController)
		topic.GET("/getById/:tid", controller.TopicGetController)
		topic.DELETE("/delete/:tid", controller.TopicDeleteController)
		topic.POST("/list", controller.TopicGetListController)
	}
	tag := router.Group("/api/tag")
	{
		tag.POST("/create", controller.TagCreateController)
		tag.POST("/update", controller.TagUpdateController)
		tag.GET("/getById/:tid", controller.TagGetController)
		tag.DELETE("/delete/:tid", controller.TagDeleteController)
		tag.POST("/list", controller.TagGetListController)
	}
	upload := router.Group("/api/upload")
	{
		upload.POST("/image", controller.UploadImageController)
	}
	return router
}
