package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	_detailKelasRepository "github.com/rg-km/final-project-engineering-4/backend/detail_kelas_siswa/repository/sqlite"
	_guruHandler "github.com/rg-km/final-project-engineering-4/backend/guru/delivery/http"
	_guruRepository "github.com/rg-km/final-project-engineering-4/backend/guru/repository/sqlite"
	_guruUseCase "github.com/rg-km/final-project-engineering-4/backend/guru/usecase"
	_kelasHandler "github.com/rg-km/final-project-engineering-4/backend/kelas/delivery/http"
	_kelasRepository "github.com/rg-km/final-project-engineering-4/backend/kelas/repository/sqlite"
	_kelasUseCase "github.com/rg-km/final-project-engineering-4/backend/kelas/usecase"
	"github.com/rg-km/final-project-engineering-4/backend/middleware"
	_orangTuaHandler "github.com/rg-km/final-project-engineering-4/backend/orangtua/delivery/http"
	_orangTuaRepository "github.com/rg-km/final-project-engineering-4/backend/orangtua/repository/sqlite"
	_orangTuaUseCase "github.com/rg-km/final-project-engineering-4/backend/orangtua/usecase"
	_siswaHandler "github.com/rg-km/final-project-engineering-4/backend/siswa/delivery/http"
	_siswaRepository "github.com/rg-km/final-project-engineering-4/backend/siswa/repository/sqlite"
	_siswaUseCase "github.com/rg-km/final-project-engineering-4/backend/siswa/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", "./kelasku.db")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middleware.CORS())

	orangTuaRepository := _orangTuaRepository.NewOrangTuaRepository(db)
	siswaRepostory := _siswaRepository.NewSiswaRepository(db)
	guruRepository := _guruRepository.NewGuruRepository(db)
	kelasRepository := _kelasRepository.NewKelasRepository(db)
	detailKelasRepository := _detailKelasRepository.NewDetailKelasSiswaRepository(db)

	orangTuaUseCase := _orangTuaUseCase.NewOrangTuaUseCase(orangTuaRepository, siswaRepostory)
	siswaUseCase := _siswaUseCase.NewSiswaUseCase(siswaRepostory)
	guruUseCase := _guruUseCase.NewGuruUseCase(guruRepository)
	kelasUseCase := _kelasUseCase.NewKelasUseCase(kelasRepository, guruRepository, siswaRepostory, orangTuaRepository, detailKelasRepository)

	_orangTuaHandler.NewOrangTuaHandler(orangTuaUseCase, r)
	_siswaHandler.NewSiswaHandler(siswaUseCase, r)
	_guruHandler.NewGuruHandler(guruUseCase, r)
	_kelasHandler.NewKelasHandler(kelasUseCase, r)

	r.Run(":8080")
}
