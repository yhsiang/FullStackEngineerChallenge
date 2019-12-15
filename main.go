package main

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/yhsiang/review360/apis"
)

func main() {
	router := gin.Default()
	secret := os.Getenv("SESSION_SECRET")
	// production use redis here
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		Secure:   true,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("review360", store))
	router.Use(apis.Database())
	api := router.Group("/api/v1")
	{
		api.POST("/signIn", apis.SignIn)

	}
	{
		basicAuth := api.Group("/")
		basicAuth.Use(apis.AuthenticationRequired())
		{
			basicAuth.GET("/employees/:id", apis.QueryEmployee)
			basicAuth.POST("/signOut", apis.SignOut)
			basicAuth.POST("/reviews", apis.CreateReview)
			basicAuth.GET("/reviews/:review_id", apis.QueryReview)
		}
	}
	adminAuth := api.Group("/admin")
	adminAuth.Use(apis.AuthenticationRequired("admin"))
	{
		adminAuth.GET("/employees", apis.QueryEmployees)
		adminAuth.POST("/employees", apis.CreateEmployee)
		adminAuth.PUT("/employees/:id", apis.UpdateEmployee)
		adminAuth.DELETE("/employees/:id", apis.RemoveEmployee)
		adminAuth.POST("/reviewers/add", apis.AddReviewer)
		adminAuth.POST("/reviewers/remove", apis.RemoveReviewer)
	}

	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}
