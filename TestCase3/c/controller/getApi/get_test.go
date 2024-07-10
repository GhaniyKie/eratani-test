package getapi

import (
	"encoding/json"
	"eratani/TestCase3/c/storage/models"
	"eratani/util/logger"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/thedevsaddam/renderer"
)

func Test_HandlerGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testFailed := []struct {
		name               string
		url                string
		wantErr            bool
		wantStatus         int
		wantMessage        string
		wantData           interface{}
		mockResultServices func(ctrl *gomock.Controller) *MockResultServices
	}{
		{
			name:        "Error input ID",
			url:         "/get",
			wantErr:     true,
			wantStatus:  http.StatusBadRequest,
			wantMessage: "Error, Bad Request",
			wantData: models.ResponseHandler{
				Message: "Error, Bad Request",
				Items:   http.StatusBadRequest,
			},
			mockResultServices: func(ctrl *gomock.Controller) *MockResultServices {
				return nil
			},
		},
		{
			name:        "Error get result services",
			url:         "/get",
			wantErr:     true,
			wantStatus:  http.StatusInternalServerError,
			wantMessage: "Something error",
			wantData: models.ResponseHandler{
				Message: "Something error",
				Items:   http.StatusInternalServerError,
			},
			mockResultServices: func(ctrl *gomock.Controller) *MockResultServices {
				return nil
			},
		},
	}

	for _, tt := range testFailed {
		t.Run(tt.name, func(t *testing.T) {
			logger := logger.New()
			render := renderer.New()

			handler := NewHandler(logger, render, tt.mockResultServices(ctrl))

			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				log.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			httpHandler := http.HandlerFunc(handler.HandlerGet)
			httpHandler.ServeHTTP(rr, req)

			// Check status
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			// Check response body
			var resBody models.ResponseHandler
			if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
				t.Fatal(err.Error())
			}
		})
	}

	testSuccess := []struct {
		name               string
		url                string
		wantErr            bool
		wantStatus         int
		wantMessage        string
		wantData           interface{}
		mockResultServices func(ctrl *gomock.Controller) *MockResultServices
	}{
		{
			name:        "Success",
			url:         "/get",
			wantErr:     false,
			wantStatus:  http.StatusOK,
			wantMessage: "Success",
			wantData: models.ResponseHandler{
				Message: "Success",
				Items:   models.ResponseServices{},
			},
			mockResultServices: func(ctrl *gomock.Controller) *MockResultServices {
				m := NewMockResultServices(ctrl)
				m.EXPECT().LogicServices(gomock.Any()).Return(models.ResponseServices{}, nil)
				return m
			},
		},
	}

	for _, tt := range testSuccess {
		t.Run(tt.name, func(t *testing.T) {
			logger := logger.New()
			render := renderer.New()

			handler := NewHandler(logger, render, tt.mockResultServices(ctrl))

			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				log.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			httpHandler := http.HandlerFunc(handler.HandlerGet)
			httpHandler.ServeHTTP(rr, req)

			// Check status
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			// Check response body
			var resBody models.ResponseHandler
			if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}
