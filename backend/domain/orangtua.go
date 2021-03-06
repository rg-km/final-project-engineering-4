package domain

//go:generate mockgen -destination=./mocks/mock_orangtua.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain OrangTuaRepository,OrangTuaUseCase

type OrangTua struct {
	ID           int64  `json:"id_wali"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHP         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	Siswa        *Siswa `json:"siswa,omitempty"`
}

type OrangTuaRepository interface {
	GetByEmail(email string) (*OrangTua, error)
	Create(orangTua OrangTua) error
}

type OrangTuaUseCase interface {
	Register(orangTua OrangTua) (*OrangTua, error)
	Login(email, password string) (*OrangTua, string, error)
}
