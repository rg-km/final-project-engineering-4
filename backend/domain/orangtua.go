package domain

//go:generate mockgen -destination=./mocks/mock_orangtua.go -package=mocks github.com/rg-km/final-project-engineering-4/backend/domain OrangTuaRepository,OrangTuaUseCase

type OrangTua struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
	Gender   string `json:"gender"`
	NoHP     string `json:"no_hp"`
	Alamat   string `json:"alamat"`
	Siswa    Siswa  `json:"siswa"`
}

type OrangTuaRepository interface {
}

type OrangTuaUseCase interface {
}
