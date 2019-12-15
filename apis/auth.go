package apis

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/yhsiang/review360/database"
	"github.com/yhsiang/review360/models"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	session := sessions.Default(c)
	user := c.PostForm("user")
	pass := c.PostForm("pass")
	if user == "admin" && pass == "admin" {
		session.Set("user", user)
		session.Set("authType", "admin")
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		c.JSON(http.StatusOK, StatusResponse{Status: true})
		return
	}

	if user == "user1" && pass == "user1" {
		var em = models.Employee{
			ID: 1,
		}
		db := c.MustGet("DB").(database.DB)
		employee, err := em.Find(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			c.Abort()
			return
		}
		session.Set("user", employee.ID)
		session.Set("authType", "employee")
		err = session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, DataResponse{Status: true, Data: employee})
		return
	}

	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Status:  false,
		Message: `invalid user and password`,
	})
}

func SignOut(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "please login first",
		})
		return
	}

	session.Clear()
	err := session.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, StatusResponse{Status: true})
}
