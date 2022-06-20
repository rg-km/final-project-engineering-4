package domain

//go:generate mockgen -destination=./mocks/mock_siswa.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain SiswaRepository,SiswaUseCase

type Siswa struct {
	ID           int64  `json:"id_siswa"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHP         string `json:"no_hp"`
	TanggalLahir string `json:"tanggal_lahir"`
	Alamat       string `json:"alamat"`
	TempatLahir  string `json:"tempat_lahir"`
	Agama        string `json:"agama"`
	Filename     string `json:"file_name"`
}

type SiswaRepository interface {
	GetByEmail(email string) (*Siswa, error)
	GetByUsername(username string) (*Siswa, error)
	Create(siswa Siswa) error
}

type SiswaUseCase interface {
	Register(siswa Siswa) (*Siswa, error)
}
