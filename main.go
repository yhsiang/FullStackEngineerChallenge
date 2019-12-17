package main

import (
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/yhsiang/review360/apis"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	secret := os.Getenv("SECRET")
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "review360",
		Key:             []byte(secret),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     apis.IdentityKey,
		PayloadFunc:     apis.PlayloadHandler,
		IdentityHandler: apis.IdentityHandler,
		Authenticator:   apis.Authenticator,
		Authorizator:    apis.Authorizator,
		Unauthorized:    apis.Unauthorized,
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.Use(apis.Database())
	router.Use(static.Serve("/", static.LocalFile("./build", false)))
	api := router.Group("/api/v1")
	{
		api.POST("/signIn", authMiddleware.LoginHandler)

	}
	{
		basicAuth := api.Group("/")
		basicAuth.Use(authMiddleware.MiddlewareFunc(), apis.AuthenticationRequired())
		{
			basicAuth.GET("/employees/:id", apis.QueryEmployee)
			basicAuth.POST("/reviews", apis.CreateReview)
			basicAuth.GET("/reviews/:review_id", apis.QueryReview)
		}
	}
	adminAuth := api.Group("/admin")
	adminAuth.Use(authMiddleware.MiddlewareFunc(), apis.AuthenticationRequired("admin"))
	{
		adminAuth.GET("/employees", apis.QueryEmployees)
		adminAuth.POST("/employees", apis.CreateEmployee)
		adminAuth.PUT("/employees/:id", apis.UpdateEmployee)
		adminAuth.DELETE("/employees/:id", apis.RemoveEmployee)
		adminAuth.POST("/reviewers/add", apis.AddReviewer)
		adminAuth.POST("/reviewers/remove", apis.RemoveReviewer)
		adminAuth.GET("/reviews", apis.QueryReviews)
		adminAuth.PUT("/reviews/:review_id", apis.UpdateReview)
	}

	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}
