package repository

import (
	"go_system_programming/cmd/specification/domain"
	"go_system_programming/cmd/specification/domainservice"
)

type userRepository struct{}

func (*userRepository) FindUserByID(id int) (*domain.User, error) {
	return &domain.User{
		ID:     1,
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		Age:    18,
		Gender: domain.GenderMale,
	}, nil
}

func NewUserRepository() domainservice.IUserRepository {
	return &userRepository{}
}
