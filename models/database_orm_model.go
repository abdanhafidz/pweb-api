package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Id                uint       `gorm:"primaryKey" json:"id"`
	UUID              uuid.UUID  `gorm:"type:uuid" json:"uuid" `
	Email             string     `gorm:"uniqueIndex" json:"email"`
	Password          string     `json:"password"`
	IsEmailVerified   bool       `json:"is_email_verified"`
	IsDetailCompleted bool       `json:"is_detail_completed"`
	CreatedAt         time.Time  `json:"created_at"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"default:null"`
}

type AccountDetails struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	AccountId   uint    `json:"account_id"`
	InitialName string  `json:"initial_name"`
	FullName    *string `json:"full_name"`
	University  *string `json:"university"`
	PhoneNumber *string `json:"phone_number"`
}

// Gorm table name settings
func (Account) TableName() string        { return "account" }
func (AccountDetails) TableName() string { return "account_details" }
