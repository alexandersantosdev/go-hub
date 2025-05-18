package model

type User struct {
	ID       int64  `gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
