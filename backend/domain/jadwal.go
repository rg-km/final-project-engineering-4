package domain

//go:generate mockgen -destination=./mocks/mock_jadwal.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain JadwalRepository,JadwalUseCase

type Jadwal struct {
	ID      int64  `json:"id"`
	Kelas   Kelas  `json:"kelas"`
	Hari    string `json:"hari"`
	Tanggal string `json:"tanggal"`
	Jam     string `json:"jam"`
	Status  string `json:"status"`
}

type JadwalRepository interface {
}

type JadwalUseCase interface {
}
