package user

import "time"

type UserFormatter struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Telepon   string `json:"telepon"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	CreatedAt time.Time
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Nama:     user.Nama,
		Password: user.Password,
		// Alamat: user.Alamat,
		// Email:  user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}

	return formatter
}
