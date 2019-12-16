package apis

import (
	"log"
	"net/http"

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

func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get(IdentityKey)

		if len(auths) != 0 && !funk.ContainsString(auths, user.(*User).Username) {
			c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
			c.Abort()
			return
		}
		c.Next()
	}
}
