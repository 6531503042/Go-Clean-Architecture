package service

import "backend/modules/user/repository"

type (
	UserService interface {
		// DefaultService
	}

	// declare userService struct
	userService struct {
		userRepository repository.UserRepository
	}
)

// declare NewUserService function
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}
