package domain

//go:generate mockgen -destination=./mocks/mock_siswa.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain SiswaRepository,SiswaUseCase

type Siswa struct {
	ID           int64  `json:"id_siswa,omitempty"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Nama         string `json:"nama,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	NoHP         string `json:"no_hp,omitempty"`
	TanggalLahir string `json:"tanggal_lahir,omitempty"`
	Alamat       string `json:"alamat,omitempty"`
	TempatLahir  string `json:"tempat_lahir,omitempty"`
	Agama        string `json:"agama,omitempty"`
	Filename     string `json:"file_name,omitempty"`
}

type SiswaRepository interface {
	GetByEmail(email string) (*Siswa, error)
	GetByUsername(username string) (*Siswa, error)
	Create(siswa Siswa) error
}

type SiswaUseCase interface {
	Register(siswa Siswa) (*Siswa, error)
	Login(username, password string) (*Siswa, string, error)
}
