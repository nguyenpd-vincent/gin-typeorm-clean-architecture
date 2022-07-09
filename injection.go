package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pdnguyen1503/base-go/handler"
	"github.com/pdnguyen1503/base-go/pkg/getenvs"
	"github.com/pdnguyen1503/base-go/repository"
	"github.com/pdnguyen1503/base-go/service"
)

func inject(d *dataSources) (*gin.Engine, error) {
	log.Println("Injecting data sources", d.DB)
	/*
	 * repository layer
	 */
	userRepo := repository.NewUserRepository(d.DB, d.MDB)

	/*
	 * service layer
	 */
	userService := service.NewUserService(&service.UserConfig{
		UserRepo: userRepo,
	})

	// read in ACCOUNT_API_URL
	baseURL := os.Getenv("ACCOUNT_API_URL")
	// read in HANDLER_TIMEOUT
	ht, err := getenvs.GetEnvInt64("HANDLER_TIMEOUT", 5)
	if err != nil {
		return nil, fmt.Errorf("could not parse HANDLER_TIMEOUT as int: %w", err)
	}

	mbb, err := getenvs.GetEnvInt64("MAX_BODY_MBS", 20)
	if err != nil {
		return nil, fmt.Errorf("could not parse MAX_BODY_BYTES as int: %w", err)
	}

	// initialize gin.Engine
	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R:               router,
		UserService:     userService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(ht) * time.Second),
		MaxBodyBytes:    mbb * 1024 * 1024,
	})

	return router, nil
}
