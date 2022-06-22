package domain

//go:generate mockgen -destination=./mocks/mock_detal_tugas.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain DetailTugasRepository,DetailTugasUseCase

type DetailTugas struct {
	ID               int64  `json:"id,omitempty"`
	Siswa            *Siswa `json:"siswa,omitempty"`
	Tugas            *Tugas `json:"tugas,omitempty"`
	WaktuPengumpulan string `json:"waktu_pengumpulan,omitempty"`
	StatusPengerjaan string `json:"status_pengerjaan,omitempty"`
}

type DetailTugasRepository interface {
}

type DetailTugasUseCase interface {
}
