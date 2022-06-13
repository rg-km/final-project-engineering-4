package domain

//go:generate mockgen -destination=./mocks/mock_siswa.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain SiswaRepository,SiswaUseCase

type Siswa struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nama         string `json:"nama"`
	Gender       string `json:"gender"`
	NoHP         string `json:"no_hp"`
	TanggalLahir string `json:"tanggal_lahir"`
	Alamat       string `json:"alamat"`
	Tempat       string `json:"tempat"`
	Filename     string `json:"filename"`
}

type SiswaRepository interface {
}

type SiswaUseCase interface {
}
