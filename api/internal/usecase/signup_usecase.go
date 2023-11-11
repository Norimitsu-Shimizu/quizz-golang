package usecase

import "app/internal/repository"

type ISignUpUsecase interface {
}

// type signUpUsecaseInteractor struct {
// 	userRepository repository.UserRepository
// }

type signUpUsecase struct {
	userRepository repository.UserRepository
}

func NewSignUpUsecase(userRepository repository.UserRepository) ISignUpUsecase {
	return &signUpUsecase{
		userRepository: userRepository,
	}
}
