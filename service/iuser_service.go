package service

import "github.com/alexandersantosdev/go-hub/dto"

type IUserService interface {
	CreateUser(user *dto.UserDTO) error
	GetUserByEmail(email string) (*dto.UserDTO, error)
	GetUserByID(ID int64) (*dto.UserDTO, error)
	UpdateUser(ID int64, user *dto.UserDTO) error
	DeleteUser(ID int64) error
}
