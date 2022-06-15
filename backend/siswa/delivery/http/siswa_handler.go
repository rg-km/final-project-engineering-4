package http

import (
	"net/http"

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
	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	// 	return
	// }

	var Siswa domain.Siswa
	err := c.BindJSON(&Siswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err = o.siswaUseCase.Register(Siswa)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
