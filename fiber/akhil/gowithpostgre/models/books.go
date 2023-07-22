package models

import "gorm.io/gorm"

type Books struct {
	ID        uint    `gorm:"primaryKey"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBoolModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
