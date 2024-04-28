package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/pkg/auth"
)

type Deps struct {
	Repos        *repository.Repository
	TokenManager auth.Manager
}

type Users interface {
	SignUp(userInput *domain.UserInput) (uint, error)
	SignIn(user *domain.User) (Tokens, error)
	RefreshTokens(refreshToken string) (Tokens, error)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Services struct {
	UsersService Users
}

func NewServices(deps Deps) *Services {
	return &Services{
		UsersService: NewUsersService(deps.Repos.UserRepo, deps.TokenManager),
	}
}
