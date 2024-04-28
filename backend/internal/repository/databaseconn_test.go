package repository_test

import (
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	if db == nil {
		t.Fatal("Failed to connect to the database")
	}

	// Правильная проверка на ошибки после выполнения запроса
	result := db.Exec("SELECT 1")
	if result.Error != nil {
		t.Fatalf("Failed to execute query on the database: %v", result.Error)
	}
}
