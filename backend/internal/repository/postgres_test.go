package repository_test

import (
	"backend/internal/domain"
	l "backend/pkg/log"
	lr "github.com/sirupsen/logrus"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	dsn := "host=localhost user=postgres password=qwerty dbname=resicomm port=5444 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.LogError("read", "Failed to connect to the database", err, lr.Fields{"err": err})
		os.Exit(1)
	}

	if err := db.Migrator().DropTable(&domain.User{}); err != nil {
		l.LogError("init", "Failed to drop user table", err, lr.Fields{"error": err})
		os.Exit(1)
	}
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		l.LogError("init", "Failed to migrate database", err, lr.Fields{"error": err})
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
