package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-4/backend/domain"
)

type guruHandler struct {
	guruUseCase domain.GuruUseCase
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewGuruHandler(guruUseCase domain.GuruUseCase, r *gin.Engine) {
	handler := &guruHandler{guruUseCase}

	guru := r.Group("/api/guru")
	{
		guru.POST("/signup", handler.Register)
		guru.POST("/login", handler.Login)
	}
}

func (g *guruHandler) Register(c *gin.Context) {

	var Guru domain.Guru
	err := c.BindJSON(&Guru)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   err.Error(),
			"data":      nil,
		})
		return
	}

	data, err := g.guruUseCase.Register(Guru)
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

func (g *guruHandler) Login(c *gin.Context) {
	var Guru domain.Guru
	err := c.BindJSON(&Guru)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"timestamp": time.Now().Format("Mon, Jan 2 2006 15:04:05"),
			"code":      http.StatusBadRequest,
			"message":   err.Error(),
			"data":      nil,
		})
		return
	}

	data, token, err := g.guruUseCase.Login(Guru.Email, Guru.Password)
	if err != nil {
		var code int
		if err == domain.ErrEmailNotFound || err == domain.ErrPasswordWrong {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}

		c.JSON(http.StatusUnauthorized, gin.H{
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
