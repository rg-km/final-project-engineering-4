package domain

//go:generate mockgen -destination=./mocks/mock_tugas.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain TugasRepository,TugasUseCase

type Tugas struct {
	ID         int64  `json:"id_tugas,omitempty"`
	Kelas      *Kelas `json:"class,omitempty"`
	JudulTugas string `json:"judul_tugas,omitempty"`
	Keterangan string `json:"keetrangan,omitempty"`
	Deadline   string `json:"deadline,omitempty"`
}

type TugasRepository interface {
}

type TugasUseCase interface {
}
