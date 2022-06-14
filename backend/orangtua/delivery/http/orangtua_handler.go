package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type orangTuaHandler struct {
	orangTuaUseCase domain.OrangTuaUseCase
}

type orangTuaRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHP         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	EmailSiswa   string `json:"email_siswa"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewOrangTuaHandler(orangTuaUseCase domain.OrangTuaUseCase, r *gin.Engine) {
	handler := &orangTuaHandler{orangTuaUseCase}

	r.POST("/api/orangtua/signup", handler.Register)
}

func (o *orangTuaHandler) Register(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var req orangTuaRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	orangTua := domain.OrangTua{
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		Nama:         req.Nama,
		JenisKelamin: req.JenisKelamin,
		NoHP:         req.NoHP,
		Alamat:       req.Alamat,
		Siswa: domain.Siswa{
			Email: req.EmailSiswa,
		},
	}

	err = o.orangTuaUseCase.Register(orangTua)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
