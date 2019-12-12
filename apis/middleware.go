package apis

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"

	"github.com/yhsiang/review360/database"
)

func Database() gin.HandlerFunc {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

// func Authenticate() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		user := c.PostForm("user")
// 		pass := c.PostForm("pass")
// 		if user == "admin" && pass == "admin" {
// 			session.Set("user", user)
// 			session.Save()
// 		}
// 		c.Next()
// 	}
// }
func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}
		if len(auths) != 0 {
			authType := session.Get("authType")
			if authType == nil || !funk.ContainsString(auths, authType.(string)) {
				c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
