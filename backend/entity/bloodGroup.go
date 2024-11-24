package entity

import "gorm.io/gorm"

type BloodGroup struct {
    gorm.Model
    Name string `json:"name"` // เช่น A, B, AB, O
	BloodGroup string `json:"blood_group"` // เช่น Rh+ หรือ Rh-
	Patients []Patient `gorm:"foreignKey:BloodGroupID"`
	Employees  []Employee `gorm:"foreignKey:BloodGroupID"` 

}