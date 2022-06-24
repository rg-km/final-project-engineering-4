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
	GetByID(id int64) (*Kelas, error)
	GetByGuruID(id int64) ([]Kelas, error)
	GetByCode(code string) (*Kelas, error)
	Create(kelas Kelas) error
}

type KelasUseCase interface {
	FetchKelas(role, email string) ([]Kelas, error)
	CreateKelas(kelas Kelas) (*Kelas, error)
	JoinKelas(emailSiswa, code string) (*Kelas, error)
}
