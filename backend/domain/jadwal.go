package domain

//go:generate mockgen -destination=./mocks/mock_jadwal.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain JadwalRepository,JadwalUseCase

type Jadwal struct {
	ID      int64  `json:"id_jadwal,omitempty"`
	Kelas   *Kelas `json:"class,omitempty"`
	Hari    string `json:"hari,omitempty"`
	Tanggal string `json:"tanggal,omitempty"`
	Jam     string `json:"jam,omitempty"`
	Status  string `json:"status,omitempty"`
}

type JadwalRepository interface {
}

type JadwalUseCase interface {
}
