package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// dbUserRepository is data/repository implement
// of service layer UserRepository
type dbUserRepository struct {
	MDB *mongo.Client
	DB  *gorm.DB
}

// NewUserRepository is a Factory for init user repository
func NewUserRepository(
	db *gorm.DB,
	mdb *mongo.Client,
) UserRepository {
	return &dbUserRepository{
		DB:  db,
		MDB: mdb,
	}
}

func (r *dbUserRepository) GetUser() (string, error) {
	return "GetUser", nil
}
