package domain

//go:generate mockgen -destination=./mocks/mock_tugas.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain TugasRepository,TugasUseCase

type Tugas struct {
	ID         int64  `json:"id"`
	Kelas      Kelas  `json:"kelas"`
	JudulTugas string `json:"judul_tugas"`
	Deadline   string `json:"deadline"`
	Keterangan string `json:"keterangan"`
}

type TugasRepository interface {
}

type TugasUseCase interface {
}
