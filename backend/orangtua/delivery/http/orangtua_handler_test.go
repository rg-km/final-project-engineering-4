package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

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

	Describe("login orangtua", func() {
		When("email given was not found", func() {
			It("should return status unauthorized", func() {
				mockOrangTuaUseCase.EXPECT().Login("wrong email", "password").Return(nil, "", domain.ErrEmailNotFound)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(handler.OrangTuaRequest{Email: "wrong email", Password: "password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/orangtua/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusUnauthorized))
			})
		})

		When("wrong password given", func() {
			It("should return status unauthorized", func() {
				mockOrangTuaUseCase.EXPECT().Login("email", "wrong password").Return(nil, "", domain.ErrPasswordWrong)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(handler.OrangTuaRequest{Email: "email", Password: "wrong password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/orangtua/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusUnauthorized))
			})
		})
		When("login success", func() {
			It("should return status OK", func() {
				mockOrangTuaUseCase.EXPECT().Login("email", "password").Return(&domain.OrangTua{Email: "email", Password: "password"}, "jwt", nil)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(handler.OrangTuaRequest{Email: "email", Password: "password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/orangtua/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
