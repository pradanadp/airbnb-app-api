package user

import "be-api/features"

type UserDataInterface interface {
	Login(email string, password string) (int, error)
	Insert(input features.UserEntity) error
	
}

type UserServiceInterface interface {
	LoginUser(email string, password string) (int, error)
	AddUser(input features.UserEntity) error
}
