package mysql

import (
	"log"
	"braw-api/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	// db.Migrator().DropTable(
	// 	&entity.User{},
	// 	&entity.Book{},
	// 	&entity.Rent{},
	// 	&entity.Role{},
	// )

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Preloved{},
		&entity.Jastip{},
		&entity.Jasantar{},
		&entity.Komunitasbraw{},
		
	); err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}