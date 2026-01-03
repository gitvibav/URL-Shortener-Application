package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitvibav/url-shortener-application/api/database"
	"github.com/gitvibav/url-shortener-application/api/models"
)

func EditURL(c *gin.Context) {
	shortID := c.Param("shortID")
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot Parse JSON",
		})
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ShortID doesn't exists",
		})
	}

	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update the shortend content",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "The Content has been updated",
	})
}
