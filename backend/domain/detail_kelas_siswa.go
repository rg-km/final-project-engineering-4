package domain

//go:generate mockgen -destination=./mocks/mock_detail_kelas_siswa.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain DetailKelasSiswaRepository,DetailKelasSiswaUseCase

type DetailKelasSiswa struct {
	ID    int64 `json:"id"`
	Siswa Siswa `json:"siswa"`
	Kelas Kelas `json:"kelas"`
}

type DetailKelasSiswaRepository interface {
}

type DetailKelasSiswaUseCase interface {
}
