package user

import "time"

type User struct {
	ID        int
	Nama      string `gorm:"column:full_name"`
	Alamat    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
