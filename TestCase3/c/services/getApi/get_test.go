package getapi

import (
	"eratani/TestCase3/c/storage/models"
	"eratani/util/logger"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_LogicServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := models.RequestData{
		Id: 1,
	}
	respForQuery := []models.ResponseQuery{}

	testCases := []struct {
		name           string
		wantErr        bool
		mockResultData func(ctrl *gomock.Controller) *MockResultData
	}{
		{
			name:    "Error",
			wantErr: true,
			mockResultData: func(ctrl *gomock.Controller) *MockResultData {
				m := NewMockResultData(ctrl)
				m.EXPECT().GetDataResponse(gomock.Any()).Return(respForQuery, errors.New("error"))
				return m
			},
		},
		{
			name:    "Success",
			wantErr: false,
			mockResultData: func(ctrl *gomock.Controller) *MockResultData {
				m := NewMockResultData(ctrl)
				m.EXPECT().GetDataResponse(gomock.Any()).Return(respForQuery, nil)
				return m
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			logger := logger.New()
			executor := NewGetApi(logger, tt.mockResultData(ctrl))

			data, err := executor.LogicServices(request)
			if !tt.wantErr && err != nil {
				t.Errorf("Error want error. got %v want %v", tt.wantErr, err)
			}
			assert.NotNil(t, data)
		})
	}
}
