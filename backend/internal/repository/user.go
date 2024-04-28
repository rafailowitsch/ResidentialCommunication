// repository/userRepo.go
package repository

import (
	"backend/internal/domain"
	l "backend/pkg/log"
	lr "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *domain.User) (uint, error) {
	l.LogInfo("create", "Attempting to create new user", lr.Fields{"email": user.Email})

	err := r.db.Create(user).Error
	if err != nil {
		l.LogError("create", "Failed to create user", err, lr.Fields{"email": user.Email})
		return 0, err
	}
	l.LogInfo("create", "User created successfully", lr.Fields{"userID": user.ID, "email": user.Email})

	return user.ID, nil
}

func (r *UserRepo) Read(id uint) (*domain.User, error) {
	l.LogInfo("read", "Attempting to read user", lr.Fields{"userID": id})

	user := &domain.User{}
	err := r.db.Preload("Company").Preload("Company.Statistics").First(user, id).Error
	if err != nil {
		l.LogError("read", "Failed to read user", err, lr.Fields{"userID": id})
		return nil, err
	}
	l.LogInfo("read", "User read successfully", lr.Fields{"userID": id})

	return user, nil
}

func (r *UserRepo) Update(user *domain.User) (uint, error) {
	l.LogInfo("update", "Attempting to update user", lr.Fields{"userID": user.ID})

	err := r.db.Save(user).Error
	if err != nil {
		l.LogError("update", "Failed to update user", err, lr.Fields{"userID": user.ID})
		return 0, err
	}

	l.LogInfo("update", "User updated successfully", lr.Fields{"userID": user.ID})
	return user.ID, nil
}

func (r *UserRepo) Delete(id uint) error {
	l.LogInfo("delete", "Attempting to delete user", lr.Fields{"userID": id})

	err := r.db.Delete(&domain.User{}, id).Error
	if err != nil {
		l.LogError("delete", "Failed to delete user", err, lr.Fields{"userID": id})
		return err
	}

	l.LogInfo("delete", "User deleted successfully", lr.Fields{"userID": id})
	return nil
}
