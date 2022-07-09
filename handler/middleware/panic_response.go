package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdnguyen1503/base-go/pkg/logging"
)

func RequestCancelRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("client cancel the request")
				logging.Error(err)

				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					map[string]string{"error": "Server internal error."},
				)
				panic(err)
			}
		}()
		c.Next()
	}
}
