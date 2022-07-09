package repository

// import (
// 	"context"
// )

type UserRepository interface {
	GetUser() (string, error)
}
