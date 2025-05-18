package repository

import (
	"github.com/alexandersantosdev/go-hub/dto"
	"gorm.io/gorm"
)

type UserPostgres struct {
	DB *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{
		DB: db,
	}
}

func CreateUser(user *dto.UserDTO) error {
	return nil
}

func GetUserByEmail(email string) (*dto.UserDTO, error) {
	return &dto.UserDTO{}, nil
}

func GetUserByID(ID int64) (*dto.UserDTO, error) {
	return &dto.UserDTO{}, nil
}

func UpdateUser(ID int64, user *dto.UserDTO) error {
	return nil
}

func DeleteUser(ID int64) error {
	return nil
}
