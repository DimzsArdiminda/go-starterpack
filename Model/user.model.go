package Model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "user"
}