package repository

import (
	"backend/internal/domain"
	l "backend/pkg/log"
	lr "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StatisticsRepo struct {
	db *gorm.DB
}

func NewStatisticsRepo(db *gorm.DB) *StatisticsRepo {
	return &StatisticsRepo{db: db}
}

func (r *StatisticsRepo) Create(statistics *domain.Statistics) (uint, error) {
	l.LogInfo("create", "Attempting to create new statistics", lr.Fields{"companyID": statistics.CompanyID})

	err := r.db.Create(statistics).Error
	if err != nil {
		l.LogError("create", "Failed to create statistics", err, lr.Fields{"companyID": statistics.CompanyID})
		return 0, err
	}
	l.LogInfo("create", "Statistics created successfully", lr.Fields{"statisticsID": statistics.ID, "companyID": statistics.CompanyID})

	return statistics.ID, nil
}

func (r *StatisticsRepo) Read(id uint) (*domain.Statistics, error) {
	l.LogInfo("read", "Attempting to read statistics", lr.Fields{"statisticsID": id})

	statistics := &domain.Statistics{}
	err := r.db.First(statistics, id).Error
	if err != nil {
		l.LogError("read", "Failed to read statistics", err, lr.Fields{"statisticsID": id})
		return nil, err
	}
	l.LogInfo("read", "Statistics read successfully", lr.Fields{"statisticsID": id})

	return statistics, nil
}

func (r *StatisticsRepo) Update(statistics *domain.Statistics) (uint, error) {
	l.LogInfo("update", "Attempting to update statistics", lr.Fields{"statisticsID": statistics.ID})

	err := r.db.Save(statistics).Error
	if err != nil {
		l.LogError("update", "Failed to update statistics", err, lr.Fields{"statisticsID": statistics.ID})
		return 0, err
	}

	l.LogInfo("update", "Statistics updated successfully", lr.Fields{"statisticsID": statistics.ID})
	return statistics.ID, nil
}

func (r *StatisticsRepo) Delete(id uint) error {
	l.LogInfo("delete", "Attempting to delete statistics", lr.Fields{"statisticsID": id})

	err := r.db.Delete(&domain.Statistics{}, id).Error
	if err != nil {
		l.LogError("delete", "Failed to delete statistics", err, lr.Fields{"statisticsID": id})
		return err
	}

	l.LogInfo("delete", "Statistics deleted successfully", lr.Fields{"statisticsID": id})
	return nil
}
