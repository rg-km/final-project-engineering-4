package domain

//go:generate mockgen -destination=./mocks/mock_presensi.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain PresensiRepository,PresensiUseCase

type Presensi struct {
	ID               int64            `json:"id"`
	Tanggal          string           `json:"tanggal"`
	DetailKelasSiswa DetailKelasSiswa `json:"detail_kelas_siswa"`
	StatusHadir      bool             `json:"status_hadir"`
}

type PresensiRepository interface {
}

type PresensiUseCase interface {
}
