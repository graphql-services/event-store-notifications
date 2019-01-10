package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Notification ...
type Notification struct {
	// gorm.Model
	IID          uuid.UUID `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	IMessage     string
	IDate        time.Time
	IPrincipal   *string
	IChannel     *string
	IReference   *string
	IReferenceID *string
	ISeen        *bool
}
