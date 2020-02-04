package handle

import (
	"net/http"

	"github.com/eastrocky/thank-you-notes-service/model"
	"github.com/eastrocky/thank-you-notes-service/repository"
	"github.com/gin-gonic/gin"
)

// ThanksGet handles requests routed to GET /thanks/:to
func ThanksGet(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		thankyous, err := repo.Get(c.Param("to"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, thankyous)
		return
	}
}

// ThanksPost handles requests routed to POST /thanks
func ThanksPost(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var thankyou model.ThankYou
		if err := c.ShouldBindJSON(&thankyou); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusText(http.StatusBadRequest),
				"message": err.Error(),
			})
			return
		}
		if err := repo.Save(thankyou); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusCreated, thankyou)
		return
	}
}
