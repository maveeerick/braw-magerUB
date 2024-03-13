package model

type CreatePreloved struct {
	Title       string `json:"title" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Condition   string `json:"condition" binding:"required"`
	Description string `json:"description" binding:"required"`
	//Stock       uint   `json:"stock" binding:"required"`
}

type UpdatePreloved struct {
	Title       string `json:"title"`
	Category    string `json:"category"`
	Price       uint   `json:"price"`
	Condition   string `json:"condition"`
	Description string `json:"description"`
	//Stock       uint   `json:"stock"`
}