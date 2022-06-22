package domain

//go:generate mockgen -destination=./mocks/mock_kelas.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain KelasRepository,KelasUseCase

type Kelas struct {
	ID          int64  `json:"id_class,omitempty"`
	Guru        *Guru  `json:"guru,omitempty"`
	NamaKelas   string `json:"nama_kelas,omitempty"`
	Keterangan  string `json:"keterangan,omitempty"`
	KodeKelas   string `json:"kode_kelas,omitempty"`
	Filename    string `json:"file_name,omitempty"`
	StatusKelas string `json:"status_kelas,omitempty"`
}

type KelasRepository interface {
}

type KelasUseCase interface {
}
