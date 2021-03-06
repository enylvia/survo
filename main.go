package main

import (
	"log"
	"path/filepath"
	"survorest/auth"
	"survorest/handler"
	"survorest/survey"
	"survorest/transactions"
	"survorest/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	webHandler "survorest/web/handler"
)

func main() {
	dsn := "survo:AVNS_pgQn9ILGvDm12asuTG2@tcp(survo-db-do-user-11081946-0.b.db.ondigitalocean.com:25060)/survo?charset=utf8mb4&parseTime=True&loc=Local"
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

	userWebHandler := webHandler.NewUserHandler(userService)
	surveyWebHandler := webHandler.NewSurveyHandler(surveyService)
	transactionWebHandler := webHandler.NewTransactionHandler(transactionService)
	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte("survosecret"))
	router.Use(sessions.Sessions("survostartup", cookieStore))
	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")

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
	api.POST("/transactionwithdraw", transactionHandler.CreateTransaction)

	router.GET("/dashboard", userWebHandler.Dashboard)
	router.POST("/session", userWebHandler.Create)
	router.Use(auth.AuthAdminMiddleware())
	router.GET("/logout", userWebHandler.Destroy)
	router.GET("/users", userWebHandler.Index)
	router.GET("/user/delete/:id", userWebHandler.Delete)
	router.GET("/surveys", surveyWebHandler.IndexSurvey)
	router.GET("/transactions", transactionWebHandler.IndexTransaction)
	router.GET("/transactions/update/:id", transactionWebHandler.UpdateTransaction)

	router.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
