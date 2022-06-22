package domain

//go:generate mockgen -destination=./mocks/mock_guru.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain GuruRepository,GuruUseCase

type Guru struct {
	ID           int64  `json:"id,omitempty"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Nama         string `json:"nama,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	NoHP         string `json:"no_hp,omitempty"`
	Pendidikan   string `json:"pendidikan,omitempty"`
	Alamat       string `json:"alamat,omitempty"`
	Agama        string `json:"agama,omitempty"`
	Filename     string `json:"filename,omitempty"`
}

type GuruRepository interface {
	GetByEmail(email string) (*Guru, error)
	Create(guru Guru) error
}

type GuruUseCase interface {
	Register(guru Guru) (*Guru, error)
	Login(username, password string) (*Guru, string, error)
}
