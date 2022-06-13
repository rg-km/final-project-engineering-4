package domain

//go:generate mockgen -destination=./mocks/mock_guru.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain GuruRepository,GuruUseCase

type Guru struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Nama       string `json:"nama"`
	Gender     string `json:"gender"`
	NoHP       string `json:"no_hp"`
	Pendidikan string `json:"pendidikan"`
	Alamat     string `json:"alamat"`
	Tempat     string `json:"tempat"`
	Filename   string `json:"filename"`
}

type GuruRepository interface {
}

type GuruUseCase interface {
}
