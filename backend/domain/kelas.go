package domain

//go:generate mockgen -destination=./mocks/mock_kelas.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain KelasRepository,KelasUseCase

type Kelas struct {
	ID          int64  `json:"id"`
	NamaKelas   string `json:"nama_kelas"`
	Keterangan  string `json:"keterangan"`
	KodeKelas   string `json:"kode_kelas"`
	StatusKelas string `json:"status_kelas"`
	Guru        Guru   `json:"guru"`
	Filename    string `json:"filename"`
}

type KelasRepository interface {
}

type KelasUseCase interface {
}
