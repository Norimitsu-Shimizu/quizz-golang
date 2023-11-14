package usecase

import (
	"app/internal/domain"
	"app/internal/parameter"
	"app/internal/repository"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type ISignUpUsecase interface {
	// Execute(c echo.Context, body io.ReadCloser)
	Execute(c echo.Context, req *parameter.PostUserRequest) error
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

func (s signUpUsecase) Execute(c echo.Context, req *parameter.PostUserRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	password := string(hashedPassword)

	newUser := domain.NewUser(req.UserName, req.Email, password)

	s.userRepository.Create(c, newUser)

	return nil
}
