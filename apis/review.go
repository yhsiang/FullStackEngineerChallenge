package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yhsiang/review360/database"
	"github.com/yhsiang/review360/models"
)

type ReviewForm struct {
	Reviewee int64  `json:"reviewee"`
	Reviewer int64  `json:"reviewer"`
	Content  string `json:"content"`
	ReviewID int64  `uri:"review_id"`
}

func QueryReview(c *gin.Context) {
	var r ReviewForm
	if err := c.ShouldBindUri(&r); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	db := c.MustGet("DB").(database.DB)
	var re = models.Review{
		ID: r.ReviewID,
	}
	review, err := re.Find(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, DataResponse{
		Status: true,
		Data:   review,
	})
}

func CreateReview(c *gin.Context) {
	var rf ReviewForm
	if err := c.ShouldBindJSON(&rf); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		c.Abort()
		return
	}
	db := c.MustGet("DB").(database.DB)
	var as = models.Assignment{
		Reviewee: rf.Reviewee,
		Reviewer: rf.Reviewer,
	}
	assignID, err := as.FindAssignID(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	var review = models.Review{
		AssignID: assignID,
		Content:  rf.Content,
	}

	_, err = review.Save(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, StatusResponse{Status: true})
}

func AddReviewer(c *gin.Context) {
	var as models.Assignment
	if err := c.ShouldBindJSON(&as); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	db := c.MustGet("DB").(database.DB)
	_, err := as.Save(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, StatusResponse{Status: true})
}

func RemoveReviewer(c *gin.Context) {
	var as models.Assignment
	if err := c.ShouldBindJSON(&as); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	db := c.MustGet("DB").(database.DB)
	err := as.Remove(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, StatusResponse{Status: true})
}
