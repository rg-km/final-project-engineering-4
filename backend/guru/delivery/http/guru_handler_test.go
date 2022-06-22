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
	handler "github.com/rg-km/final-project-engineering-4/backend/guru/delivery/http"
)

var _ = Describe("GuruHandler", func() {
	var (
		ctrl            *gomock.Controller
		mockGuruUseCase *mocks.MockGuruUseCase
		r               *gin.Engine
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockGuruUseCase = mocks.NewMockGuruUseCase(ctrl)
		r = gin.Default()
		handler.NewGuruHandler(mockGuruUseCase, r)
	})

	AfterEach(func() {
		defer ctrl.Finish()
	})

	Describe("login guru", func() {
		When("email given was not found", func() {
			It("should return status unauthorized", func() {
				mockGuruUseCase.EXPECT().Login("wrong email", "password").Return(nil, "", domain.ErrEmailNotFound)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(domain.Guru{Email: "wrong email", Password: "password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/guru/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusUnauthorized))
			})
		})

		When("wrong password given", func() {
			It("should return status unauthorized", func() {
				mockGuruUseCase.EXPECT().Login("email", "wrong password").Return(nil, "", domain.ErrPasswordWrong)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(domain.Guru{Email: "email", Password: "wrong password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/guru/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusUnauthorized))
			})
		})
		When("login success", func() {
			It("should return status OK", func() {
				mockGuruUseCase.EXPECT().Login("email", "password").Return(&domain.Guru{Email: "email", Password: "password"}, "jwt", nil)
				var body bytes.Buffer
				json.NewEncoder(&body).Encode(domain.Guru{Email: "email", Password: "password"})

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/guru/login", &body)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
