package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	UserRepo     repository.User
	TokenManager auth.Manager
}

func NewUsersService(userRepo repository.User, tokenManager auth.Manager) *UsersService {
	return &UsersService{UserRepo: userRepo, TokenManager: tokenManager}
}

func (s *UsersService) SignUp(userInput *domain.UserInput) (uint, error) {
	var id uint

	bPasswordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	if err != nil {
		return 0, err
	}
	passwordHash := string(bPasswordHash)

	user := domain.User{
		Email:        userInput.Email,
		PasswordHash: passwordHash,
		FirstName:    userInput.FirstName,
		LastName:     userInput.LastName,
		Age:          userInput.Age,
		Address:      userInput.Address,
		UserType:     domain.Resident,
	}

	if id, err = s.UserRepo.Create(&user); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *UsersService) SignIn(user *domain.User) (Tokens, error) {
	return Tokens{}, nil
}

func (s *UsersService) RefreshTokens(refreshToken string) (Tokens, error) {
	return Tokens{}, nil
}
