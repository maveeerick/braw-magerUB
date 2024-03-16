package entity

import (
	"time"

	"github.com/google/uuid"
)

type Jastip struct {
	ID_Jastip   uuid.UUID	 `json:"id_jastip" gorm:"type:varchar(36);primary_key;"`
	ID_User     uuid.UUID    `json:"id_user" gorm:"type:varchar(36);foreign_key;"`
	Title      	string   	 `json:"title" gorm:"type:varchar(255);not null;"`
	Category  	string		 `json:"category" gorm:"type:varchar(255);not null;"`	
	Price 		int			 `json:"price" gorm:"type:integer;not null;"`
	Open_Day	string		 `json:"open_day" gorm:"type:varchar(255);not null;"`
	Close_Order time.Time	 	 `json:"close_order" gorm:"autoClose_Order"`	
	PhotoLink 	string    	 `json:"-" gorm:"type:varchar(200)"`
	CreatedAt 	time.Time 	 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt 	time.Time 	 `json:"updatedAt" gorm:"autoUpdateTime"`
}