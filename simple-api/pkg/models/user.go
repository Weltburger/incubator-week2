package models

type User struct {
	ID        string `json:"id" xml:"id" gorm:"primaryKey"`
	Email     string `json:"email" xml:"email"`
	Password  string `json:"password" xml:"password"`
}
