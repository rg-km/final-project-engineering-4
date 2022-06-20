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

	r.POST("/api/siswa/signup", handler.Register)
}

func (o *siswaHandler) Register(c *gin.Context) {

	var Siswa domain.Siswa
	err := c.BindJSON(&Siswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   err.Error(),
			"data":      nil,
		})
		return
	}

	data, err := o.siswaUseCase.Register(Siswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
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
