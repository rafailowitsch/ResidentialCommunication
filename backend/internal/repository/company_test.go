package repository_test

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCompanyRepo_Create(t *testing.T) {
	repo := repository.NewCompanyRepo(db)

	company := &domain.Company{
		Name:    "Test Company" + time.Now().Format("2006010215040515"),
		Address: "123 Test St",
		City:    "TestCity",
	}

	id, err := repo.Create(company)
	if err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	assert.NotZero(t, id)
}

func TestCompanyRepo_Read(t *testing.T) {
	repo := repository.NewCompanyRepo(db)

	company := &domain.Company{
		Name:    "Test Company" + time.Now().Format("2006010215040515"),
		Address: "123 Test St",
		City:    "TestCity",
	}

	id, err := repo.Create(company)
	if err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	readCompany, err := repo.Read(id)
	assert.NoError(t, err)
	assert.Equal(t, company.Name, readCompany.Name)
}

func TestCompanyRepo_Update(t *testing.T) {
	repo := repository.NewCompanyRepo(db)

	company := &domain.Company{
		Name:    "Test Company" + time.Now().Format("2006010215040515"),
		Address: "123 Test St",
		City:    "TestCity",
	}

	id, err := repo.Create(company)
	if err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	company.ID = id
	company.Name = "Updated Company Name"
	_, err = repo.Update(company)
	assert.NoError(t, err)

	updatedCompany, err := repo.Read(id)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Company Name", updatedCompany.Name)
}

func TestCompanyRepo_Delete(t *testing.T) {
	repo := repository.NewCompanyRepo(db)

	company := &domain.Company{
		Name:    "Test Company" + time.Now().Format("2006010215040515"),
		Address: "123 Test St",
		City:    "TestCity",
	}

	id, err := repo.Create(company)
	if err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	err = repo.Delete(id)
	assert.NoError(t, err)

	_, err = repo.Read(id)
	assert.Error(t, err)
}
