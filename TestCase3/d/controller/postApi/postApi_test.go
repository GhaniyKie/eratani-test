package postapi

import (
	"encoding/json"
	models "eratani/TestCase3/d/storage/models"
	"eratani/util/logger"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/thedevsaddam/renderer"
)

func Test_HandlerPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := models.ResponseData{}

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
			name:        "Error decode body",
			url:         "/post",
			wantErr:     true,
			wantStatus:  http.StatusInternalServerError,
			wantMessage: "Error decode body",
			wantData: models.ResponseData{
				Message: "Error decode body",
				Data:    http.StatusInternalServerError,
			},
			mockResultServices: func(ctrl *gomock.Controller) *MockResultServices {
				m := NewMockResultServices(ctrl)
				m.EXPECT().LogicServices(gomock.Any()).Return(resp, errors.New("Error decode body"))
				return m
			},
		},
		{
			name:        "Error param input",
			url:         "/post",
			wantErr:     true,
			wantStatus:  http.StatusBadRequest,
			wantMessage: "Error, column cant nil",
			wantData: models.ResponseData{
				Message: "Error, column cant nil",
				Data:    http.StatusBadRequest,
			},
			mockResultServices: func(ctrl *gomock.Controller) *MockResultServices {
				m := NewMockResultServices(ctrl)
				m.EXPECT().LogicServices(gomock.Any()).Return(resp, errors.New("Error, column cant nil"))
				return m
			},
		},
		{
			name:        "Error get data services",
			url:         "/post",
			wantErr:     true,
			wantStatus:  http.StatusInternalServerError,
			wantMessage: "Error something",
			wantData: models.ResponseData{
				Message: "Error something",
				Data:    http.StatusInternalServerError,
			},
			mockResultServices: func(ctrl *gomock.Controller) *MockResultServices {
				m := NewMockResultServices(ctrl)
				m.EXPECT().LogicServices(gomock.Any()).Return(resp, errors.New("Errorc column cant nil"))
				return m
			},
		},
	}

	for _, tt := range testFailed {
		t.Run(tt.name, func(t *testing.T) {
			logger := logger.New()
			render := renderer.New()

			handler := NewHandler(logger, render, tt.mockResultServices(ctrl))

			req, err := http.NewRequest(http.MethodPost, tt.url, nil)
			if err != nil {
				log.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			httpHandler := http.HandlerFunc(handler.HandlerPost)
			httpHandler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			var resBody models.ResponseData
			if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
				t.Fatal(err.Error())
			}
		})

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
				url:         "/post",
				wantErr:     false,
				wantStatus:  http.StatusOK,
				wantMessage: "Success",
				wantData: models.ResponseData{
					Message: "Success",
					Data:    http.StatusOK,
				},
				mockResultServices: func(ctrl *gomock.Controller) *MockResultServices {
					m := NewMockResultServices(ctrl)
					m.EXPECT().LogicServices(gomock.Any()).Return(resp, nil)
					return m
				},
			},
		}

		for _, tt := range testSuccess {
			t.Run(tt.name, func(t *testing.T) {
				logger := logger.New()
				render := renderer.New()

				handler := NewHandler(logger, render, tt.mockResultServices(ctrl))

				req, err := http.NewRequest(http.MethodPost, tt.url, nil)
				if err != nil {
					log.Fatalf(err.Error())
				}

				rr := httptest.NewRecorder()
				httpHandler := http.HandlerFunc(handler.HandlerPost)
				httpHandler.ServeHTTP(rr, req)

				if status := rr.Code; status != tt.wantStatus {
					t.Errorf("handler returned wrong status code: got %v want %v",
						status, tt.wantStatus)
				}

				var resBody models.ResponseData
				if err := json.NewDecoder(rr.Body).Decode(&resBody); err != nil {
					t.Fatal(err.Error())
				}
			})
		}
	}
}
