package service

import (
	"github.com/alexandersantosdev/go-hub/dto"
	"github.com/alexandersantosdev/go-hub/repository"
)

type UserServiceImpl struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

// CreateUser implements UserService.
func (u *UserServiceImpl) CreateUser(user *dto.UserDTO) error {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (u *UserServiceImpl) DeleteUser(ID int64) error {
	panic("unimplemented")
}

// GetUserByEmail implements UserService.
func (u *UserServiceImpl) GetUserByEmail(email string) (*dto.UserDTO, error) {
	panic("unimplemented")
}

// GetUserByID implements UserService.
func (u *UserServiceImpl) GetUserByID(ID int64) (*dto.UserDTO, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *UserServiceImpl) UpdateUser(ID int64, user *dto.UserDTO) error {
	panic("unimplemented")
}
