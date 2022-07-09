package main

import (
	// "context"
	"log"
	"net/http"

	// "time"

	"github.com/pdnguyen1503/base-go/pkg/logging"
)

func init() {
	logging.Setup()

}

// @title Auth Service API
// @version 1.0
// @description Service auth of hxb
// @termsOfService https://github.com/pdnguyen1503/base-go
// @license.name MIT
func main() {
	log.Println("Starting server...")

	//initalizing data_source
	ds, err := initDS()

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}
	logging.Info("ds", ds)

	router, err := inject(ds)

	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}
	srv := &http.Server{
		Addr:    ":3001",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()
	log.Printf("Listening on port %v\n", srv.Addr)

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// // shutdown data sources
	// if err := ds.close(); err != nil {
	// 	log.Fatalf("A problem occurred gracefully shutting down data sources: %v\n", err)
	// }

	// // Shutdown server
	// log.Println("Shutting down server...")
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Server forced to shutdown: %v\n", err)
	// }
}
