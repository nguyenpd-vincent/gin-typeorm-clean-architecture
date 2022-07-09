package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdnguyen1503/base-go/pkg/app"
	"github.com/pdnguyen1503/base-go/pkg/e"
	"github.com/pdnguyen1503/base-go/pkg/logging"
)

// Create payment
// @Summary CreatePayment
// @Produce  json
// @Tags User
// @Success 200 {object} object
// @Failure 401,400,500 {object} app.ResponseError
// @Router / [GET]
func (h *Handler) GetUser(c *gin.Context) {
	var (
		// req  dto.CreatePaymentDto
		appG = app.Gin{C: c}
	)
	userId := c.GetString("uId")
	if userId == "" {
		err := e.NewBadRequest("No user provinid")
		appG.ResponseError(err, nil)
		return
	}
	// if ok := request.BindData(c, &req); !ok {
	// 	return
	// }

	p, err := h.UserService.GetUser()
	if err != nil || p == "" {
		logging.Error(err)
		err := e.NewBadRequest("Can not create payment")
		appG.ResponseError(err, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"data":    p,
	})
}
