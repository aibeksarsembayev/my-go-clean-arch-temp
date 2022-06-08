package usecase

import (
	"context"

	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type userUsecase struct {
	userRepository models.UserRepository
}

// NewUserUsecase will create new an user Usecase object representation of models.UserUsecase interface
func NewUserUsecase(u models.UserRepository) models.UserUsecase {
	return &userUsecase{
		userRepository: u,
	}
}

// Create user ...
func (u *userUsecase) Create(ctx context.Context, user *models.User) (id int, err error) {
	return 0, nil
}

// Update user ...
func (u *userUsecase) Update(ctx context.Context, user *models.User) (err error) {
	return nil
}

// GetByID user ...
func (u *userUsecase) GetByID(ctx context.Context, id int) (user *models.User, err error) {
	return nil, nil
}

// GetByEmail user ...
func (u *userUsecase) GetByEmail(ctx context.Context, email string) (user *models.User, err error) {
	return nil, nil
}

// Delete user ...
func (u *userUsecase) Delete(ctx context.Context, id int) (err error) {
	return nil
}
