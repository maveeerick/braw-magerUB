package mysql

import (
	"log"
	//"math/rand"

	"braw-api/entity"
	//"github.com/bxcodec/faker/v4"
	//"github.com/google/uuid"
	"gorm.io/gorm"
)



func generateRole(db *gorm.DB) error {
	var roles []*entity.Role

	roles = append(roles,
		&entity.Role{
			ID:   1,
			Role: "admin",
		},
		&entity.Role{
			ID:   2,
			Role: "user",
		})

	if err := db.CreateInBatches(roles, 2).Error; err != nil {
		return err
	}
	return nil
}

func SeedData(db *gorm.DB) {
	var totalRole int64
	if err := db.Model(&entity.Role{}).Count(&totalRole).Error; err != nil {
		log.Fatalf("Error while counting book: %v", err)
	}

	if totalRole == 0 {
		if err := generateRole(db); err != nil {
			log.Fatalf("Error while generating book: %v", err)

		}
	}
}