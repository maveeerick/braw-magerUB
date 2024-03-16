package entity

import(
	"github.com/google/uuid"
	"time"
)

type Preloved struct {
	ID_Preloved uuid.UUID	 `json:"id_preloved" gorm:"type:varchar(36);primary_key;"`
	ID_User     uuid.UUID    `json:"id_user" gorm:"type:varchar(36);foreign_key;"`
	Title      	string   	 `json:"title" gorm:"type:varchar(255);not null;"`
	Category  	string		 `json:"category" gorm:"type:varchar(255);not null;"`	
	Price  	  	int		 	 `json:"price" gorm:"type:integer;not null;"`
	Condition 	string		 `json:"condition" gorm:"type:varchar(255);not null"`
	Description	string		 `json:"description" gorm:"type:varchar(255)"`
	PhotoLink 	string    	 `json:"-" gorm:"type:varchar(200)"`
	CreatedAt 	time.Time 	 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt 	time.Time 	 `json:"updatedAt" gorm:"autoUpdateTime"`
}