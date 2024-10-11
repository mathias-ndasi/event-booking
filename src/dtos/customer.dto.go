package dtos

type SignUpDto struct {
	EmailAddress string `json:"emailAddress" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

type LoginDto struct {
	SignUpDto
}
