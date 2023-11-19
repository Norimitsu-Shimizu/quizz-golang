package usecase

import (
	"app/internal/dto"
	"app/internal/parameter"
	"app/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ISignInUsecase interface {
	Execute(c echo.Context, req *parameter.PostLoginRequest) (*dto.UserResponse, error)
}

type signInUsecase struct {
	userRepo repository.UserRepository
}

type jwtCustomClaims struct {
	ID uint
	jwt.StandardClaims
}

func NewSignInUsecase(userRepository repository.UserRepository) ISignInUsecase {
	return &signInUsecase{
		userRepo: userRepository,
	}
}

func (s *signInUsecase) Execute(c echo.Context, req *parameter.PostLoginRequest) (*dto.UserResponse, error) {
	// メールアドレスからユーザーを取得
	user, err := s.userRepo.GetByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}

	// JWTを生成
	token, err := generateJWTToken(user.ID)
	if err != nil {
		return nil, err
	}

	return dto.ToUserResponse(user, token), nil
}

func generateJWTToken(userID uint) (string, error) {
	// JWTを生成
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
