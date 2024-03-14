package mysql

import (
	"log"
	"math/rand"

	"braw-api/entity"
	//"github.com/bxcodec/faker/v4"
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

func generatePreloved(db *gorm.DB) error {
	categorys := []string{"Sci-Fi", "Fantasy", "Mystery", "Romance", "History"}
	var preloveds []*entity.Preloved

	for i := 1; i <= 20; i++ {
		preloved := &entity.Preloved{
			// ID:          uuid.New(),
			// Title:       faker.Sentence(),
			// Writter:     faker.Name(),
			// Year:        rand.Intn(20) + 2000,
			 Category:       categorys[rand.Intn(len(categorys))],
			// Description: faker.Paragraph(),
			// Stock:       10,
		}
		preloveds = append(preloveds, preloved)
	}

	if err := db.CreateInBatches(preloveds, 20).Error; err != nil {
		return err
	}
	return nil
}


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

	var totalPreloved int64
	if err := db.Model(&entity.Preloved{}).Count(&totalPreloved).Error; err != nil {
		log.Fatalf("Error while counting preloved: %v", err)
	}

	if totalPreloved == 0 {
		if err := generatePreloved(db); err != nil {
			log.Fatalf("Error while generating preloved: %v", err)

		}
	}

	var totalRole int64
	if err := db.Model(&entity.Role{}).Count(&totalRole).Error; err != nil {
		log.Fatalf("Error while counting preloved: %v", err)
	}

	if totalRole == 0 {
		if err := generateRole(db); err != nil {
			log.Fatalf("Error while generating preloved: %v", err)

		}
	}
}