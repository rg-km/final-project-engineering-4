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
	ID          int64  `json:"id_presensi,omitempty"`
	IDSiswa     int64  `json:"id_siswa,omitempty"`
	IDKelas     int64  `json:"id_kelas,omitempty"`
	Tanggal     string `json:"tanggal,omitempty"`
	Jam         string `json:"jam,omitempty"`
	StatusHadir string `json:"status_hadir"`
}

func NewPresensiHandler(presensiUseCase domain.PresensiUseCase, r *gin.Engine) {
	handler := &presensiHandler{presensiUseCase}

	presensi := r.Group("/api/presensi")
	{
		presensi.GET("/:id", middleware.Auth("Orang Tua"), handler.FetchPresensiSiswa)
		presensi.POST("/create", middleware.Auth("Guru"), handler.CreatePresensiSiswa)
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
	res, err := p.presensiUseCase.FetchPresensiSiswa(int64(id), email)
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

func (p *presensiHandler) CreatePresensiSiswa(c *gin.Context) {

	var req PresensiRequest
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

	presensi := domain.Presensi{
		Jam:         req.Jam,
		Tanggal:     req.Tanggal,
		StatusHadir: req.StatusHadir,
	}
	res, err := p.presensiUseCase.CreatePresensiSiswa(req.IDSiswa, req.IDKelas, presensi)

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
		"message":   "Presensi Berhasil Dibuat",
		"data":      res,
	})

}
