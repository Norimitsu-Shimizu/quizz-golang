package usecase

import "app/internal/repository"

// type SignUpUsecase interface {
// }

// type signUpUsecaseInteractor struct {
// 	userRepository repository.UserRepository
// }

type SignUpUsecase struct {
	userRepository repository.UserRepository
}

func NewSignUpUsecase(userRepository repository.UserRepository) *SignUpUsecase {
	return &SignUpUsecase{
		userRepository: userRepository,
	}
}
