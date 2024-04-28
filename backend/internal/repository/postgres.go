package repository

import (
	"backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.Company{}, &domain.User{}, &domain.Appeal{}, &domain.Feedback{}, &domain.Statistics{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
