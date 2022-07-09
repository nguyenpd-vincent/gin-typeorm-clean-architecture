package service

type UserService interface {
	GetUser() (string, error)
}
