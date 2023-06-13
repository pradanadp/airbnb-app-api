package user

import "be-api/features"

type UserDataInterface interface {
	Login(email string, password string) (int, error)
	Insert(input features.UserEntity) error
	Select(id int) (features.UserEntity,error)
	Delete(id int) error
	SelectId(id int) error
	Update(input features.UserEntity,id uint) error
	
}

type UserServiceInterface interface {
	LoginUser(email string, password string) (int, error)
	AddUser(input features.UserEntity) error
	GetUser(id int) (features.UserEntity,error)
	DeleteUser(id int) error
	GetId(id int) error
	Update(input features.UserEntity,id uint) error
}
