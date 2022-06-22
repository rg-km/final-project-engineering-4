package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	_guruHandler "github.com/rg-km/final-project-engineering-4/backend/guru/delivery/http"
	_guruRepository "github.com/rg-km/final-project-engineering-4/backend/guru/repository/sqlite"
	_guruUseCase "github.com/rg-km/final-project-engineering-4/backend/guru/usecase"
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

	orangTuaRepository := _orangTuaRepository.NewOrangTuaRepository(db)
	siswaRepostory := _siswaRepository.NewSiswaRepository(db)
	guruRepository := _guruRepository.NewGuruRepository(db)

	orangTuaUseCase := _orangTuaUseCase.NewOrangTuaUseCase(orangTuaRepository, siswaRepostory)
	siswaUseCase := _siswaUseCase.NewSiswaUseCase(siswaRepostory)
	guruUseCase := _guruUseCase.NewGuruUseCase(guruRepository)

	_orangTuaHandler.NewOrangTuaHandler(orangTuaUseCase, r)
	_siswaHandler.NewSiswaHandler(siswaUseCase, r)
	_guruHandler.NewGuruHandler(guruUseCase, r)

	r.Run(":8080")
}
