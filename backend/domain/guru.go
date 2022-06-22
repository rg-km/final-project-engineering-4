package domain

//go:generate mockgen -destination=./mocks/mock_guru.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain GuruRepository,GuruUseCase

type Guru struct {
	ID           int64  `json:"id_guru,omitempty"`
	Password     string `json:"password,omitempty"`
	Nama         string `json:"nama,omitempty"`
	Alamat       string `json:"alamat,omitempty"`
	NoHP         string `json:"no_hp,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	Agama        string `json:"agama,omitempty"`
	Email        string `json:"email,omitempty"`
	Pendidikan   string `json:"pendidikan,omitempty"`
	Filename     string `json:"file_name,omitempty"`
}

type GuruRepository interface {
}

type GuruUseCase interface {
}
