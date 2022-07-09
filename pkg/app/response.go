package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pdnguyen1503/base-go/pkg/e"
	"github.com/pdnguyen1503/base-go/pkg/logging"
)

type Gin struct {
	C *gin.Context
}

type ResponseError struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) ResponseError(err error, data interface{}) {
	logging.Error("Error Response", err)
	g.C.JSON(e.Status(err), ResponseError{
		Error: err.Error(),
		Data:  data,
	})
	return
}
