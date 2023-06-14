package data

import (
	"be-api/features"
	"be-api/features/user"
	"be-api/utils"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

// Upgrade implements user.UserDataInterface
func (repo *UserData) Upgrade(input features.UserEntity, id uint) (error) {
	var user features.User
	tx := repo.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return tx.Error
	}

	updatedUser := features.UserEntityToModel(input)
	if updatedUser.NamaPemilik == "" && updatedUser.KTPFile == "" && updatedUser.NIBFile==""{
		return errors.New("data tidak boleh kosong")
	}
	updatedUser.Role = "hoster"
	updateOpr := repo.db.Model(&user).Updates(updatedUser)
	if updateOpr.Error != nil {
		return errors.New(updateOpr.Error.Error() + ", failed to update user")
	}

	return nil
}

// Update implements user.UserDataInterface
func (repo *UserData) Update(input features.UserEntity, id uint) (error) {
	var user features.User
	tx := repo.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return tx.Error
	}

	updatedUser := features.UserEntityToModel(input)
	updateOpr := repo.db.Model(&user).Updates(updatedUser)
	if updateOpr.Error != nil {
		return errors.New(updateOpr.Error.Error() + ", failed to update user")
	}
	return nil
}

// SelectId implements user.UserDataInterface
func (repo *UserData) SelectId(id int) error {
	var user features.User
	tx := repo.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements user.UserDataInterface
func (repo *UserData) Delete(id int) error {
	tx := repo.db.Delete(&features.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Select implements user.UserDataInterface
func (repo *UserData) Select(id int) (features.UserEntity, error) {
	var user features.User
	tx := repo.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return features.UserEntity{}, tx.Error
	}

	mapUser := features.UserModelToEntity(user)

	return mapUser, nil
}

// Insert implements user.UserDataInterface
func (repo *UserData) Insert(input features.UserEntity) (uint,error) {
	hashPassword, err := utils.HashPasword(input.Password)
	if err != nil {
		return 0,errors.New("error hashing password: " + err.Error())
	}
	input.Password = hashPassword
	userData := features.UserEntityToModel(input)

	if userData.Username == "" {
		str := strconv.Itoa(15)
		userData.Username = userData.FirstName + "." + userData.LastName + str

		tx := repo.db.Create(&userData)
		if tx.Error != nil {
			return 0,tx.Error
		} else if tx.RowsAffected == 0 {
			return 0,errors.New("insert data user failed, rows affected 0 ")
		}
		id := userData.ID
		return id,nil
	}

	tx := repo.db.Create(&userData)
	if tx.Error != nil {
		return 0,tx.Error
	} else if tx.RowsAffected == 0 {
		return 0,errors.New("insert data user failed, rows affected 0 ")
	}
	id := userData.ID
	return id,nil
}

// Login implements user.UserDataInterface
func (repo *UserData) Login(email string, password string) (int, error) {
	var user features.User
	tx := repo.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return 0, errors.New("email tidak terdaftar")
	}
	match := utils.CheckPaswordHash(password, user.Password)
	if !match {
		return 0, errors.New("password tidak cocok")
	}
	return int(user.ID), nil
}

func New(db *gorm.DB) user.UserDataInterface {
	return &UserData{
		db: db,
	}
}
