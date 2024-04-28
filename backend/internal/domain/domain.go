package domain

import (
	"gorm.io/gorm"
	"time"
)

// Company представляет управляющую компанию
type Company struct {
	gorm.Model
	Name         string      `gorm:"size:255;not null;uniqueIndex"`
	Address      string      `gorm:"size:255;not null"`
	City         string      `gorm:"size:100;not null"`
	Users        []User      `gorm:"foreignKey:CompanyID"`
	Appeals      []Appeal    `gorm:"foreignKey:CompanyID"`
	StatisticsID *uint       // Изменено на *uint
	Statistics   *Statistics `gorm:"foreignKey:StatisticsID"`
}

// User представляет собой структуру для хранения информации о пользователе.
type User struct {
	gorm.Model
	Email        string `gorm:"size:100;uniqueIndex;not null"`
	PasswordHash string `gorm:"size:255;not null"`

	FirstName  string `gorm:"size:255;not null"`
	MiddleName string `gorm:"size:255"`
	LastName   string `gorm:"size:255;not null"`
	Age        int    `gorm:"default:0"`
	Address    string `gorm:"size:255"`

	UserType UserType
	//CompanyID uint    `gorm:"not null;index"`
	//Company   Company `gorm:"foreignKey:CompanyID"`
}

type Session struct {
	gorm.Model
	RefreshToken string    `gorm:"size:255;uniqueIndex;not null"`
	ExpiresAt    time.Time `gorm:"not null"`
}

// UserType представляет тип пользователя.
type UserType int

const (
	Representative UserType = iota // Представитель управляющей компании
	Resident                       // Жилец дома
)

// Appeal представляет собой структуру для хранения информации о обращении пользователя.
type Appeal struct {
	gorm.Model
	Type        AppealType `gorm:"not null"`
	Subject     string     `gorm:"size:255;not null"`
	Description string     `gorm:"size:1000;not null"`
	Date        time.Time  `gorm:"not null"`
	//Files       []string   `gorm:"type:json"`
	Status AppealStatus

	UserID     uint     `gorm:"not null;index"`
	CompanyID  uint     `gorm:"not null;index"`
	FeedbackID uint     `gorm:"unique"`
	Feedback   Feedback `gorm:"foreignKey:FeedbackID"`
}

// AppealType представляет тип обращения.
type AppealType int

const (
	Complaint AppealType = iota
	RepairRequest
	Suggestion
	InformationRequest
	Emergency
	UtilityIssue
)

// AppealStatus представляет статус обращения.
type AppealStatus int

const (
	New AppealStatus = iota
	InProgress
	Completed
	Rejected
	RequiresAdditionalInfo
)

// Feedback представляет собой структуру для хранения информации об обратной связи.
type Feedback struct {
	gorm.Model
	AppealID uint `gorm:"not null;unique;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Rating      int       `gorm:"type:integer;not null"`
	Comment     string    `gorm:"size:1000"`
	IsAnonymous bool      `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
}

// Statistics представляет собой структуру для хранения обобщенной статистики.
type Statistics struct {
	gorm.Model
	CompanyID uint `gorm:"not null;unique;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	AppealsCount      int `gorm:"not null"`
	AppealsAssessment int `gorm:"not null"`
	//AvgResponseTime    time.Duration `gorm:"not null"`
}

type UserInput struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
}
