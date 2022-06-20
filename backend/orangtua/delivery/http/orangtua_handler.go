package http

import (
	"net/http"
	"time"

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

func NewOrangTuaHandler(orangTuaUseCase domain.OrangTuaUseCase, r *gin.Engine) {
	handler := &orangTuaHandler{orangTuaUseCase}

	r.POST("/api/orangtua/signup", handler.Register)
}

func (o *orangTuaHandler) Register(c *gin.Context) {
	var req orangTuaRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   err.Error(),
			"data":      nil,
		})
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

	data, err := o.orangTuaUseCase.Register(orangTua)
	if err != nil {
		var code int
		if err == domain.ErrEmailSiswaNotFound {
			code = http.StatusNotFound
		} else {
			code = http.StatusBadRequest
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      code,
			"message":   err.Error(),
			"data":      nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
		"code":      http.StatusOK,
		"message":   "Registrasi berhasil",
		"data":      data,
	})
}
