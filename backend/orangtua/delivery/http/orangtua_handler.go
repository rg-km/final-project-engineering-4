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

type OrangTuaRequest struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Nama         string `json:"nama,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	NoHP         string `json:"no_hp,omitempty"`
	Alamat       string `json:"alamat,omitempty"`
	EmailSiswa   string `json:"email_siswa,omitempty"`
}

func NewOrangTuaHandler(orangTuaUseCase domain.OrangTuaUseCase, r *gin.Engine) {
	handler := &orangTuaHandler{orangTuaUseCase}

	orangtua := r.Group("/api/orangtua")
	{
		orangtua.POST("/signup", handler.Register)
		orangtua.POST("/login", handler.Login)
	}
}

func (o *orangTuaHandler) Register(c *gin.Context) {
	var req OrangTuaRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   domain.ErrInvalidRequest.Error(),
			"data":      nil,
		})
		return
	}

	orangTua := domain.OrangTua{
		Email:        req.Email,
		Password:     req.Password,
		Nama:         req.Nama,
		JenisKelamin: req.JenisKelamin,
		NoHP:         req.NoHP,
		Alamat:       req.Alamat,
		Siswa: &domain.Siswa{
			Email: req.EmailSiswa,
		},
	}

	data, err := o.orangTuaUseCase.Register(orangTua)
	if err != nil {
		var code int
		if err == domain.ErrEmailSiswaNotFound {
			code = http.StatusNotFound
		} else {
			code = http.StatusInternalServerError
		}

		c.JSON(code, gin.H{
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

func (o *orangTuaHandler) Login(c *gin.Context) {
	var req OrangTuaRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   domain.ErrInvalidRequest.Error(),
			"data":      nil,
		})
		return
	}

	data, token, err := o.orangTuaUseCase.Login(req.Email, req.Password)
	if err != nil {
		var code int
		if err == domain.ErrEmailNotFound {
			code = http.StatusNotFound
		} else if err == domain.ErrPasswordWrong {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}

		c.JSON(code, gin.H{
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
		"message":   "Login Berhasil",
		"data": gin.H{
			"token": token,
			"user":  data,
		},
	})
}
