package domain

//go:generate mockgen -destination=./mocks/mock_detail_kelas_siswa.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain DetailKelasSiswaRepository,DetailKelasSiswaUseCase

type DetailKelasSiswa struct {
	ID    int64  `json:"id_dt_kelassiswa,omitempty"`
	Siswa *Siswa `json:"siswa,omitempty"`
	Kelas *Kelas `json:"class,omitempty"`
}

type DetailKelasSiswaRepository interface {
	GetBySiswaID(siswaID int64) ([]DetailKelasSiswa, error)
	GetByKelasID(kelasID int64) ([]DetailKelasSiswa, error)
	Create(detailKelas DetailKelasSiswa) error
}

type DetailKelasSiswaUseCase interface {
}
