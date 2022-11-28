package user

type RegisterUserInput struct {
	Nama     string `json:"nama" `
	Alamat   string `json:"alamat" `
	Telepon  string `json:"telepon" `
	Email    string `json:"email"`
	Password string `json:"password" `
}

type LoginInput struct {
	Email    string `json:"nama"`
	Password string `json:"password"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type FormCreateUserInput struct {
	Nama            string `json:"nama" binding:"required"`
	Alamat          string `json:"alamat" binding:"required"`
	Telepon         string `json:"telepon" binding:"required"`
	JenisKelamin    string `json:"jenis_kelamin" binding:"required"`
	Role            string `json:"role" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	GoogleSignId    string `json:"google_sign_id" binding:"required"`
	SyaratKetentuan string `json:"syarat_ketentuan" binding:"required"`
	Error           error
}

type FormUpdateUserInput struct {
	ID           int
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	Telepon      string `json:"telepon"`
	JenisKelamin string `json:"jenis_kelamin"`
	Email        string `json:"email"`
	Password     string `json:"password" binding:"required"`
	Error        error
}
