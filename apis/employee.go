package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yhsiang/review360/database"
	"github.com/yhsiang/review360/models"
)

func QueryEmployees(c *gin.Context) {
	db := c.MustGet("DB").(database.DB)
	var em models.Employee
	employees, err := em.FindAll(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, DataResponse{
		Status: true,
		Data:   employees,
	})
}

func CreateEmployee(c *gin.Context) {
	var em models.Employee
	if err := c.ShouldBindJSON(&em); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	db := c.MustGet("DB").(database.DB)
	employee, err := em.Save(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, DataResponse{
		Status: true,
		Data:   employee,
	})
}

func UpdateEmployee(c *gin.Context) {
	var em models.Employee
	if err := c.ShouldBindUri(&em); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&em); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	db := c.MustGet("DB").(database.DB)
	employee, err := em.Save(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, DataResponse{
		Status: true,
		Data:   employee,
	})
}

func QueryEmployee(c *gin.Context) {
	var em models.Employee
	if err := c.ShouldBindUri(&em); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	db := c.MustGet("DB").(database.DB)
	employee, err := em.Find(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, DataResponse{
		Status: true,
		Data:   employee,
	})
}

func RemoveEmployee(c *gin.Context) {
	var em models.Employee
	if err := c.ShouldBindUri(&em); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	db := c.MustGet("DB").(database.DB)
	err := em.Remove(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, StatusResponse{Status: true})
}
