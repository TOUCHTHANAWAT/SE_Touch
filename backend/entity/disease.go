package entity

import (
	"gorm.io/gorm"
)

type Disease struct {
	gorm.Model
	DiseaseName string     `json:"disease_name"`
	Employees   []Employee `gorm:"many2many:employee_diseases;" json:"employees"` // Many-to-Many relationship with Employee

	Patients   []Patient `gorm:"many2many:patient_diseases;" json:"patients"` // Many-to-Many relationship with Employee

	MedicalRecords []MedicalRecords `gorm:"many2many:medicalrecords_diseases;" json:"medicalrecords"`
}