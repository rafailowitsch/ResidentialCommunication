package repository

import (
	"backend/internal/domain"
)

type Company interface {
	Create(company *domain.Company) (uint, error)
	Read(id uint) (*domain.Company, error)
	Update(company *domain.Company) (uint, error)
	Delete(id uint) error
}

type User interface {
	Create(user *domain.User) (uint, error)
	Read(id uint) (*domain.User, error)
	Update(user *domain.User) (uint, error)
	Delete(id uint) error
}

type Appeal interface {
	Create(appeal *domain.Appeal) (uint, error)
	Read(id uint) (*domain.Appeal, error)
	Update(appeal *domain.Appeal) (uint, error)
	Delete(id uint) error
}

type Feedback interface {
	Create(feedback *domain.Feedback) (uint, error)
	Read(id uint) (*domain.Feedback, error)
	Update(feedback *domain.Feedback) (uint, error)
	Delete(id uint) error
}

type Statistics interface {
	Create(statistics *domain.Statistics) (uint, error)
	Read(id uint) (*domain.Statistics, error)
	Update(statistics *domain.Statistics) (uint, error)
	Delete(id uint) error
}

type Repository struct {
	UserRepo User
}

func NewRepository(userRepo User) *Repository {
	return &Repository{UserRepo: userRepo}
}
