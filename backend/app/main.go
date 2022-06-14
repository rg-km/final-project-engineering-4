package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	_orangTuaHandler "github.com/rg-km/final-project-engineering-4/backend/orangtua/delivery/http"
	_orangTuaRepository "github.com/rg-km/final-project-engineering-4/backend/orangtua/repository/sqlite"
	_orangTuaUseCase "github.com/rg-km/final-project-engineering-4/backend/orangtua/usecase"
	_siswaRepository "github.com/rg-km/final-project-engineering-4/backend/siswa/repository/sqlite"
)

func main() {
	db, err := sql.Open("sqlite3", "./kelasku.db")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	orangTuaRepository := _orangTuaRepository.NewOrangTuaRepository(db)
	siswaRepostory := _siswaRepository.NewSiswaRepository(db)

	orangTuaUseCase := _orangTuaUseCase.NewOrangTuaUseCase(orangTuaRepository, siswaRepostory)

	_orangTuaHandler.NewOrangTuaHandler(orangTuaUseCase, r)

	r.Run(":8080")
}
