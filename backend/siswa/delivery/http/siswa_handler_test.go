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
	handler "github.com/rg-km/final-project-engineering-4/backend/siswa/delivery/http"
)

var _ = Describe("SiswaHandler", func() {
	var (
		ctrl             *gomock.Controller
		mockSiswaUseCase *mocks.MockSiswaUseCase
		r                *gin.Engine
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockSiswaUseCase = mocks.NewMockSiswaUseCase(ctrl)
		r = gin.Default()
		handler.NewSiswaHandler(mockSiswaUseCase, r)
	})

	AfterEach(func() {
		defer ctrl.Finish()
	})

	Describe("login siswa", func() {
		When("email given was not found", func() {
			It("should return status unauthorized", func() {
				mockSiswaUseCase.EXPECT().Login("wrong email", "password").Return(nil, "", domain.ErrEmailNotFound)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(domain.Siswa{Email: "wrong email", Password: "password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/siswa/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusUnauthorized))
			})
		})

		When("wrong password given", func() {
			It("should return status unauthorized", func() {
				mockSiswaUseCase.EXPECT().Login("email", "wrong password").Return(nil, "", domain.ErrPasswordWrong)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(domain.Siswa{Email: "email", Password: "wrong password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/siswa/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusUnauthorized))
			})
		})
		When("login success", func() {
			It("should return status OK", func() {
				mockSiswaUseCase.EXPECT().Login("email", "password").Return(&domain.Siswa{Email: "email", Password: "password"}, "jwt", nil)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(domain.Siswa{Email: "email", Password: "password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/siswa/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
