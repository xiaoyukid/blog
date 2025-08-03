package main

import (
	"blog/middleware"
	"blog/repositories"
	"blog/service"
	"blog/web"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	// 初始化数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 初始化 repository
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)

	// 初始化 service
	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)

	// 初始化 web handler
	userHandler := web.NewUserHandler(userService)
	postHandler := web.NewPostHandler(postService)

	// 设置路由
	r := gin.Default()

	public := r.Group("/")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	authorized := r.Use(middleware.AuthorMiddleware())
	{
		authorized.POST("/createPosts", postHandler.CreatePost)
		authorized.POST("/post/getById", postHandler.GetPostById)
		authorized.POST("/post/getList", postHandler.GetPosts)

		authorized.POST("/posts/update", postHandler.UpdatePost)
		authorized.POST("/posts/delete", postHandler.DeletePost)

	}

	// 启动服务器
	r.Run(":8080")
}
