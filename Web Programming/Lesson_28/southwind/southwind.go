package Southwind

import (
	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string `gorm:"not null;size:30"`
	LastName  string `gorm:"not null;size:30"`
	Emails    []Email
}

type Email struct {
	gorm.Model
	EmployeeID int    `gorm:"index"`
	Mail       string `gorm:"type:varchar(50);unique_index"`
	IsActive   bool
}
