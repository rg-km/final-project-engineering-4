package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/middleware"
)

type kelasHandler struct {
	kelasUseCase domain.KelasUseCase
}

func NewKelasHandler(kelasUseCase domain.KelasUseCase, r *gin.Engine) {
	handler := &kelasHandler{kelasUseCase}

	kelas := r.Group("/api/kelas")
	{
		kelas.GET("/", middleware.Auth("all"), handler.FetchKelasSiswa)
		kelas.GET("/:id", middleware.Auth("all"), handler.FetchKelasByID)
		kelas.POST("/create", middleware.Auth("Guru"), handler.CreateKelas)
		kelas.POST("/join", middleware.Auth("Siswa"), handler.JoinKelas)
	}
}

func (k *kelasHandler) FetchKelasSiswa(c *gin.Context) {
	email := c.GetString("Email")
	role := c.GetString("Role")

	res, err := k.kelasUseCase.FetchKelas(role, email)
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

func (k *kelasHandler) CreateKelas(c *gin.Context) {
	emailGuru := c.GetString("Email")

	var req domain.Kelas
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

	req.Guru = &domain.Guru{
		Email: emailGuru,
	}
	res, err := k.kelasUseCase.CreateKelas(req)
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
		"message":   "Kelas Berhasil Dibuat",
		"data":      res,
	})
}

func (k *kelasHandler) FetchKelasByID(c *gin.Context) {
	email := c.GetString("Email")
	role := c.GetString("Role")
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

	kelas, err := k.kelasUseCase.FetchKelasByID(role, email, int64(id))
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
		"data":      kelas,
	})
}

func (k *kelasHandler) JoinKelas(c *gin.Context) {
	emailSiswa := c.GetString("Email")

	var req domain.Kelas
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

	res, err := k.kelasUseCase.JoinKelas(emailSiswa, req.KodeKelas)
	if err != nil {
		var code int
		if err == domain.ErrEmailNotFound || err == domain.ErrKelasNotFound {
			code = http.StatusNotFound
		} else if err == domain.ErrClassJoined {
			code = http.StatusBadRequest
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
		"message":   "Berhasil Join Kelas",
		"data":      res,
	})
}
