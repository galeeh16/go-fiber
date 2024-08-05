package entity

import "time"

// kalo pake * artinya bisa null balikannya, kalo gadikasih * suka ada isinya
type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Address   string     `json:"address"`
	Phone     string     `json:"phone"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at"`
}
