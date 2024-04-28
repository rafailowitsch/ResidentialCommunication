package repository_test

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func uniqueName(prefix string) string {
	return prefix + time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(1000))
}

func TestUserRepo_Create(t *testing.T) {
	repo := repository.NewUserRepo(db)

	companyName := uniqueName("Test Company")
	company := &domain.Company{
		Name:    companyName,
		Address: "123 Test St",
		City:    "TestCity",
	}
	if err := db.Create(company).Error; err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	userEmail := "test@example.com_" + time.Now().Format("20060102150405")
	user := &domain.User{
		Email:     userEmail,
		Password:  "password",
		FirstName: "Test",
		LastName:  "User",
		CompanyID: company.ID,
	}

	id, err := repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func TestUserRepo_Read(t *testing.T) {
	repo := repository.NewUserRepo(db)

	companyName := uniqueName("Test Company")
	company := &domain.Company{
		Name:    companyName,
		Address: "123 Test St",
		City:    "TestCity",
	}
	if err := db.Create(company).Error; err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	userEmail := uniqueName("test@example.com_")
	user := &domain.User{
		Email:     userEmail,
		Password:  "password",
		FirstName: "Test",
		LastName:  "User",
		CompanyID: company.ID,
	}

	id, err := repo.Create(user)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	readUser, err := repo.Read(id)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, readUser.Email)
}

func TestUserRepo_Update(t *testing.T) {
	repo := repository.NewUserRepo(db)

	company := &domain.Company{
		Name:    "Test Company" + time.Now().Format("20060102150405"),
		Address: "123 Test St",
		City:    "TestCity",
	}
	if err := db.Create(company).Error; err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	user := &domain.User{
		Email:     "test@example.com",
		Password:  "password",
		FirstName: "Test",
		LastName:  "User",
		CompanyID: company.ID,
	}

	id, err := repo.Create(user)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	user.ID = id
	user.Email = "updated@example.com"
	_, err = repo.Update(user)
	assert.NoError(t, err)

	updatedUser, err := repo.Read(id)
	assert.NoError(t, err)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

func TestUserRepo_Delete(t *testing.T) {
	repo := repository.NewUserRepo(db)

	companyName := uniqueName("Test Company")
	company := &domain.Company{
		Name:    companyName,
		Address: "123 Test St",
		City:    "TestCity",
	}
	if err := db.Create(company).Error; err != nil {
		t.Fatalf("Failed to create test company: %v", err)
	}

	userEmail := uniqueName("test@example.com_")
	user := &domain.User{
		Email:     userEmail,
		Password:  "password",
		FirstName: "Test",
		LastName:  "User",
		CompanyID: company.ID,
	}

	id, err := repo.Create(user)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	err = repo.Delete(id)
	assert.NoError(t, err)

	_, err = repo.Read(id)
	assert.Error(t, err)
}
