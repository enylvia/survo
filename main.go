package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"survorest/auth"
	"survorest/handler"
	"survorest/user"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/survo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())

	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	googleHandler := auth.GoogleService(userRepository)
	router := gin.Default()
	//router.Use(cors.Default())

	// static image route
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/oauthlogin", auth.HandleLogin)
	api.GET("/callback", googleHandler.HandleCallback)

	//Change grouping
	api.Use(auth.AuthMiddleware(userService)) // protect all routes
		api.PUT("/update/:id", userHandler.UpdateProfile)
		api.PUT("/upload/:id", userHandler.UploadAvatar)
		api.GET("/profile/:id", userHandler.GetProfile)

	router.Run(":8080")
}
