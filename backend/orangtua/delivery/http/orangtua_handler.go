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

func NewOrangTuaHandler(orangTuaUseCase domain.OrangTuaUseCase, r *gin.Engine) {
	handler := &orangTuaHandler{orangTuaUseCase}

	r.POST("/api/orangtua/signup", handler.Register)
}

func (o *orangTuaHandler) Register(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var orangTua domain.OrangTua
	err = json.Unmarshal(body, &orangTua)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = o.orangTuaUseCase.Register(orangTua)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
