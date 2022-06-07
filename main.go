package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"survorest/auth"
	"survorest/handler"
	"survorest/survey"
	"survorest/transactions"
	"survorest/user"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/survo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())

	}
	userRepository := user.NewRepository(db)
	surveyRepository := survey.NewRepository(db)
	transactionRepository := transactions.NewRepository(db)

	userService := user.NewService(userRepository)
	surveyService := survey.NewService(surveyRepository, userRepository)
	transactionService := transactions.NewService(transactionRepository)

	userHandler := handler.NewUserHandler(userService)
	surveyHandler := handler.NewSurveyHandler(surveyService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
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
	api.GET("/surveylist", surveyHandler.SurveyList)
	api.GET("/surveylist/:id", surveyHandler.SurveyList)
	api.GET("/surveydetail/:id", surveyHandler.GetSurveyDetail)
	api.GET("/surveyrespond/:id", surveyHandler.GetRespondSurvey)
	//Change grouping
	api.Use(auth.AuthMiddleware(userService)) // protect all routes
	api.GET("/transactions", transactionHandler.GetAllTransaction)

	api.PUT("/update/:id", userHandler.UpdateProfile)
	api.PUT("/upload/:id", userHandler.UploadAvatar)
	api.GET("/profile/:id", userHandler.GetProfile)
	api.POST("/createsurvey", surveyHandler.CreateSurvey)
	api.POST("/answerquestion", surveyHandler.AnswerQuestion)
	api.GET("/transaction/:id", transactionHandler.GetAllTransactionByIDUser)
	api.POST("/transactionpremium", transactionHandler.CreateTransactionPremium)
	api.POST("/transactionwithdraw",transactionHandler.CreateTransaction)

	router.Run(":8080")
}
