package entity

type Role struct {
	Id   uint   `json:"idRole" gorm:"primary_key;autoIncrement;"`
	Role string `json:"role" gorm:"type:varchar(255);not null;unique;"`
}
