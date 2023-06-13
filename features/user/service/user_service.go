package service

import (
	"be-api/features"
	"be-api/features/user"
	"fmt"

	"github.com/go-playground/validator"
)

type UserService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// UpgradeUser implements user.UserServiceInterface
func (service *UserService) UpgradeUser(input features.UserEntity, id uint) error {
	err := service.userData.Upgrade(input, id)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil	
}

// Update implements user.UserServiceInterface
func (service *UserService) Update(input features.UserEntity, id uint) error {
	err := service.userData.Update(input, id)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

// GetId implements user.UserServiceInterface
func (service *UserService) GetId(id int) error {
	err := service.userData.SelectId(id)
	if err != nil {
		return err
	}
	return nil
}

// SelectId implements user.UserServiceInterface

// DeleteUser implements user.UserServiceInterface
func (service *UserService) DeleteUser(id int) error {
	err := service.userData.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetUser implements user.UserServiceInterface
func (service *UserService) GetUser(id int) (features.UserEntity, error) {
	dataUser, err := service.userData.Select(id)
	if err != nil {
		return features.UserEntity{}, err
	}
	return dataUser, err
}

// AddUser implements user.UserServiceInterface
func (service *UserService) AddUser(input features.UserEntity) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	err := service.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

// LoginUser implements user.UserServiceInterface
func (service *UserService) LoginUser(email string, password string) (int, error) {
	loginInput := features.LoginUser{
		Email:    email,
		Password: password,
	}
	errValidate := service.validate.Struct(loginInput)
	if errValidate != nil {
		fmt.Println(errValidate)
		return 0, errValidate
	}
	userId, err := service.userData.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func New(userData user.UserDataInterface) user.UserServiceInterface {
	return &UserService{
		userData: userData,
		validate: validator.New(),
	}
}
