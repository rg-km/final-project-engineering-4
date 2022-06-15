package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type siswaHandler struct {
	siswaUseCase domain.SiswaUseCase
}

type siswaRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHP         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Agama        string `json:"agama"`
	Filename     string `json:"file_name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewSiswaHandler(siswaUseCase domain.SiswaUseCase, r *gin.Engine) {
	handler := &siswaHandler{siswaUseCase}

	r.POST("/api/siswa/signup", handler.Register)
}

func (o *siswaHandler) Register(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var req siswaRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	siswa := domain.Siswa{
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		Nama:         req.Nama,
		JenisKelamin: req.JenisKelamin,
		NoHP:         req.NoHP,
		Alamat:       req.Alamat,
		TempatLahir:  req.TempatLahir,
		TanggalLahir: req.TanggalLahir,
		Agama:        req.Agama,
		Filename:     req.Filename,
	}

	err = o.siswaUseCase.Register(siswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
