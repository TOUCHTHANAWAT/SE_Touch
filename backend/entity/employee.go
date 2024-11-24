package entity

import (
	"time"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName           string            `json:"first_name"`
	LastName            string            `json:"last_name"`
	DateOfBirth         time.Time         `json:"date_of_birth"`
	Email               string            `json:"email"`
	Phone               string            `json:"phone"`
	Address             string            `json:"address"`
	Username            string            `json:"username"`
	ProfessionalLicense string            `json:"professional_license"`
	Graduate            string            `json:"graduate"`
	NationalID          string            `json:"national_id"`
	InfoConfirm         bool              `json:"info_confirm"`


	GenderID   uint   `json:"gender_id"`
	Gender     Gender `gorm:"foreignKey:GenderID"`

	PositionID uint     `json:"position_id"`
	Position   Position `gorm:"foreignKey:PositionID"`

	DepartmentID uint       `json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID"`

	StatusID uint   `json:"status_id"`
	Status   Status `gorm:"foreignKey:StatusID"`

	SpecialistID uint       `json:"specialist_id"`
	Specialist   Specialist `gorm:"foreignKey:SpecialistID"`

	// Many-to-Many relationship with Disease
	Diseases []Disease `gorm:"many2many:employee_diseases;" json:"diseases"`

	Profile  string `json:"profile" gorm:"type:longtext"`
	Password string `json:"password"`

	TakeAHistorys []TakeAHistory `gorm:"foreignKey:EmployeeID"`

	BloodGroupID uint        `json:"blood_group_id"`
	BloodGroup   BloodGroup  `gorm:"foreignKey:BloodGroupID"`

	// ฟิลด์สำหรับ Reset Token
	ResetToken       string    `gorm:"size:36" json:"reset_token"`         // UUID ขนาด 36 ตัว
	ResetTokenExpiry time.Time `json:"reset_token_expiry"`  

	MedicalRecords []MedicalRecords `gorm:"foreignKey:EmployeeID"`
}