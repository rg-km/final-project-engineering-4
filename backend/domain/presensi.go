package domain

//go:generate mockgen -destination=./mocks/mock_presensi.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain PresensiRepository,PresensiUseCase

type Presensi struct {
	ID               int64             `json:"id_presensi,omitempty"`
	DetailKelasSiswa *DetailKelasSiswa `json:"dt_kelassiswa,omitempty"`
	Tanggal          string            `json:"tanggal,omitempty"`
	Jam              string            `json:"jam,omitempty"`
	StatusHadir      string            `json:"status_hadir"`
}

type PresensiRepository interface {
}

type PresensiUseCase interface {
}
