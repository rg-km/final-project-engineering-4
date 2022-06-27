package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/middleware"
)

type presensiHandler struct {
	presensiUseCase domain.PresensiUseCase
}

type PresensiRequest struct {
	ID string `json:"id_presensi,omitempty"`
	// Siswa       string `json:"id_siswa,omitempty"`
	Tanggal     string `json:"tanggal,omitempty"`
	Jam         string `json:"jam,omitempty"`
	StatusHadir string `json:"status_hadir,omitempty"`
}

func NewPresensiHandler(presensiUseCase domain.PresensiUseCase, r *gin.Engine) {
	handler := &presensiHandler{presensiUseCase}

	presensi := r.Group("/api/presensi")
	{
		presensi.GET("/", middleware.Auth("Orang Tua"), handler.FetchPresensiSiswa)
	}
}

func (p *presensiHandler) FetchPresensiSiswa(c *gin.Context) {
	email := c.GetString("Email")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   domain.ErrInvalidRequest.Error(),
			"data":      nil,
		})
		return
	}
	res, err := p.presensiUseCase.FetchPresensi(int64(id), email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusNotFound,
			"message":   err.Error(),
			"data":      nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
		"code":      http.StatusOK,
		"data":      res,
	})
}
