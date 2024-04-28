package repository

import (
	"backend/internal/domain"
	"backend/pkg/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CompanyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) *CompanyRepo {
	return &CompanyRepo{db: db}
}

func (r *CompanyRepo) Create(company *domain.Company) (uint, error) {
	log.LogInfo("create", "Attempting to create new company", logrus.Fields{"name": company.Name})

	err := r.db.Create(company).Error
	if err != nil {
		log.LogError("create", "Failed to create company", err, logrus.Fields{"name": company.Name})
		return 0, err
	}
	log.LogInfo("create", "Company created successfully", logrus.Fields{"companyID": company.ID, "name": company.Name})

	return company.ID, nil
}

func (r *CompanyRepo) Read(id uint) (*domain.Company, error) {
	log.LogInfo("read", "Attempting to read company", logrus.Fields{"companyID": id})

	company := &domain.Company{}
	err := r.db.First(company, id).Error
	if err != nil {
		log.LogError("read", "Failed to read company", err, logrus.Fields{"companyID": id})
		return nil, err
	}
	log.LogInfo("read", "Company read successfully", logrus.Fields{"companyID": id})

	return company, nil
}

func (r *CompanyRepo) Update(company *domain.Company) (uint, error) {
	log.LogInfo("update", "Attempting to update company", logrus.Fields{"companyID": company.ID})

	err := r.db.Save(company).Error
	if err != nil {
		log.LogError("update", "Failed to update company", err, logrus.Fields{"companyID": company.ID})
		return 0, err
	}

	log.LogInfo("update", "Company updated successfully", logrus.Fields{"companyID": company.ID})
	return company.ID, nil
}

func (r *CompanyRepo) Delete(id uint) error {
	log.LogInfo("delete", "Attempting to delete company", logrus.Fields{"companyID": id})

	err := r.db.Delete(&domain.Company{}, id).Error
	if err != nil {
		log.LogError("delete", "Failed to delete company", err, logrus.Fields{"companyID": id})
		return err
	}

	log.LogInfo("delete", "Company deleted successfully", logrus.Fields{"companyID": id})
	return nil
}
