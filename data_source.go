package main

import (
	"errors"
	"fmt"
	"log"

	// "github.com/jinzhu/gorm"
	"github.com/pdnguyen1503/base-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// _ "github.com/go-sql-driver/mysql"

type dataSources struct {
	DB  *gorm.DB
	MDB *mongo.Client
}

func initDS() (*dataSources, error) {
	log.Printf("initializing data source\n")
	config, err := configs.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	fmt.Println("a ...any", config)
	var db *gorm.DB
	db = configs.ConnectMySql(config)

	var mdb *mongo.Client
	configs.StartMongo(config)
	mdb = configs.GetMongoClient()
	if mdb == nil {
		return nil, errors.New("Can not connect to mongodb")
	}
	return &dataSources{
		DB:  db,
		MDB: mdb,
	}, nil

}

// close to be used in graceful server shutdown
func (d *dataSources) close() error {
	// sqlDB, err := d.DB
	if err := d.DB; err != nil {
		return fmt.Errorf("error closing mysql: %w", err)
	}

	return nil
}
