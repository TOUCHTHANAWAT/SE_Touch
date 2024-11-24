package entity

import (
    "gorm.io/gorm"
    "time"
)

type Patient struct {
    gorm.Model
    IdentificationNumber string  `json:"identification_number"`
    FirstName       string       `json:"first_name"`
    LastName        string       `json:"last_name"`
    DateOfBirth     time.Time    `json:"date_of_birth"`
    Address         string       `json:"address"`          
    PhoneNumber     string       `json:"phone_number"`     


    // PatientDiseases []PatientDisease `gorm:"foreignKey:PatientID" `
    Diseases []Disease 			`gorm:"many2many:patient_diseases;" json:"diseases"`
    
	GenderID        uint   		`json:"gender_id"`
    Gender     		Gender     	`gorm:"foreignKey:GenderID"`
	
    BloodGroupID    uint        `json:"blood_group_id"`
    BloodGroup 		BloodGroup 	`gorm:"foreignKey:BloodGroupID"`

    TakeAHistorys  []TakeAHistory `gorm:"foreignKey:PatientID"` 
}
