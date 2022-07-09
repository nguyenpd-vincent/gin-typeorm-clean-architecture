package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pdnguyen1503/base-go/handler/middleware"
	"github.com/pdnguyen1503/base-go/service"
)

type Handler struct {
	UserService service.UserService
}

type Config struct {
	R               *gin.Engine
	UserService     service.UserService
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService: c.UserService,
	}

	g := c.R.Group(c.BaseURL)
	g.Use(middleware.RequestCancelRecover())
	fmt.Println("gin.Mode()", gin.Mode())
	if gin.Mode() != gin.TestMode {
		// g.GET("/payment", middleware.AuthUser(h.TokenService), h.)
		g.GET("/", h.GetUser)
	} else {
		// g.GET("/me", h.Me)
	}

	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// g.GET("/health_check", httpcheck.GinHealthCheck)
}
