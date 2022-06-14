package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rg-km/final-project-engineering-4/backend/domain"
	"github.com/rg-km/final-project-engineering-4/backend/domain/mocks"
	handler "github.com/rg-km/final-project-engineering-4/backend/orangtua/delivery/http"
)

var _ = Describe("OrangtuaHandler", func() {
	var (
		ctrl                *gomock.Controller
		mockOrangTuaUseCase *mocks.MockOrangTuaUseCase
		r                   *gin.Engine
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockOrangTuaUseCase = mocks.NewMockOrangTuaUseCase(ctrl)
		r = gin.Default()
		handler.NewOrangTuaHandler(mockOrangTuaUseCase, r)

	})

	AfterEach(func() {
		defer ctrl.Finish()
	})

	Describe("Register", func() {
		var (
			mockOrangTua     domain.OrangTua
			mockOrangTuaJson []byte
		)

		BeforeEach(func() {
			mockOrangTua = domain.OrangTua{
				Username: "Username1",
				Password: "Password1",
				Nama:     "Nama1",
				Alamat:   "Alamat1",
				NoHP:     "NoHP1",
				Email:    "Email1",
				Siswa: domain.Siswa{
					Email: "EmailSiswa1",
				},
			}
			mockOrangTuaJson, _ = json.Marshal(mockOrangTua)
		})

		When("register success", func() {
			It("should return 200", func() {
				mockOrangTuaUseCase.EXPECT().Register(gomock.Any()).Return(nil)
				reqBody := strings.NewReader(string(mockOrangTuaJson))
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/orangtua/signup", reqBody)
				r.ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
