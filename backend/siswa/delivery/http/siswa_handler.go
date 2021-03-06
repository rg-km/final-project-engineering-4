package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type siswaHandler struct {
	siswaUseCase domain.SiswaUseCase
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewSiswaHandler(siswaUseCase domain.SiswaUseCase, r *gin.Engine) {
	handler := &siswaHandler{siswaUseCase}

	siswa := r.Group("/api/siswa")
	{
		siswa.POST("/signup", handler.Register)
		siswa.POST("/login", handler.Login)
	}

}

func (s *siswaHandler) Register(c *gin.Context) {

	var Siswa domain.Siswa
	err := c.BindJSON(&Siswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   domain.ErrInvalidRequest.Error(),
			"data":      nil,
		})
		return
	}

	data, err := s.siswaUseCase.Register(Siswa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusInternalServerError,
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

func (s *siswaHandler) Login(c *gin.Context) {
	var Siswa domain.Siswa
	err := c.BindJSON(&Siswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   domain.ErrInvalidRequest,
			"data":      nil,
		})
		return
	}

	data, token, err := s.siswaUseCase.Login(Siswa.Email, Siswa.Password)
	if err != nil {
		var code int
		if err == domain.ErrEmailNotFound || err == domain.ErrPasswordWrong {
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
		"message":   "Login berhasil",
		"data": gin.H{
			"token": token,
			"user":  data,
		},
	})
}
