package postapi

import (
	"eratani/TestCase3/d/storage/models"
	"eratani/util/logger"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_LogicServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	request := models.RequestPost{
		Country:        "Indonesia",
		CreditCardType: "mastercard",
		CreditCard:     12345,
		FirstName:      "firstname",
		LastName:       "lastname",
	}
	resp := models.ResponsePost{}

	testCases := []struct {
		name                  string
		wantErr               bool
		mockResultQueryInsert func(ctrl *gomock.Controller) *MockResultQueryInsert
	}{
		{
			name:    "Error get data from mysql",
			wantErr: true,
			mockResultQueryInsert: func(ctrl *gomock.Controller) *MockResultQueryInsert {
				m := NewMockResultQueryInsert(ctrl)
				m.EXPECT().PostData(request).Return(resp, errors.New("error get result query insert data"))
				return m
			},
		},
		{
			name:    "Success",
			wantErr: false,
			mockResultQueryInsert: func(ctrl *gomock.Controller) *MockResultQueryInsert {
				m := NewMockResultQueryInsert(ctrl)
				m.EXPECT().PostData(request).Return(resp, nil)
				return m
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			logger := logger.New()
			executors := NewExecutor(logger, tt.mockResultQueryInsert(ctrl))

			data, err := executors.LogicServices(request)
			if !tt.wantErr && err != nil {
				t.Errorf("Error want error. got %v want %v", tt.wantErr, err)
			}
			assert.NotNil(t, data)
		})
	}
}
