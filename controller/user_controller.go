package controller

import (
	"github.com/alexandersantosdev/go-hub/dto"
	"github.com/alexandersantosdev/go-hub/service"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(user *dto.UserDTO) error {
	return u.UserService.CreateUser(user)
}

func (u *UserController) GetUserByEmail(email string) (*dto.UserDTO, error) {
	return u.UserService.GetUserByEmail(email)
}

func (u *UserController) GetUserByID(ID int64) (*dto.UserDTO, error) {
	return u.UserService.GetUserByID(ID)
}

func (u *UserController) UpdateUser(ID int64, user *dto.UserDTO) error {
	return u.UserService.UpdateUser(ID, user)
}

func (u *UserController) DeleteUser(ID int64) error {
	return u.UserService.DeleteUser(ID)
}
