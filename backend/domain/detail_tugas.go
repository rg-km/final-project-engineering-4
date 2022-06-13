package domain

//go:generate mockgen -destination=./mocks/mock_detal_tugas.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain DetailTugasRepository,DetailTugasUseCase

type DetailTugas struct {
	ID               int64  `json:"id"`
	Siswa            Siswa  `json:"siswa"`
	Tugas            Tugas  `json:"tugas"`
	WaktuPengumpulan string `json:"waktu_pengumpulan"`
	StatusPengerjaan bool   `json:"status_pengerjaan"`
}

type DetailTugasRepository interface {
}

type DetailTugasUseCase interface {
}
