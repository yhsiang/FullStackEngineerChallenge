package apis

import (
	"fmt"
	"regexp"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/yhsiang/review360/database"
	"github.com/yhsiang/review360/models"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"user" binding:"required"`
	Password string `form:"pass" binding:"required"`
}

type User struct {
	Username string
}

const IdentityKey = "id"

func PlayloadHandler(data interface{}) jwt.MapClaims {
	if v, ok := data.(*User); ok {
		return jwt.MapClaims{
			IdentityKey: v.Username,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &User{
		Username: claims[IdentityKey].(string),
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	fmt.Printf("testtt")
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	if userID == "admin" && password == "admin" {
		return &User{
			Username: userID,
		}, nil
	}

	re := regexp.MustCompile(`^user(\d+)`)
	if re.Match([]byte(userID)) && userID == password {
		matchID := re.FindSubmatch([]byte(userID))
		employeeID, err := strconv.ParseInt(string(matchID[1]), 10, 64)
		if err != nil {
			return nil, err
		}
		em := models.Employee{
			ID: employeeID,
		}
		db := c.MustGet("DB").(database.DB)
		_, err = em.Find(db)
		if err != nil {
			return nil, err
		}
		return &User{
			Username: userID,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	re := regexp.MustCompile(`^user(\d+)`)
	if v, ok := data.(*User); ok && (v.Username == "admin" || re.Match([]byte(v.Username))) {
		return true
	}

	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
