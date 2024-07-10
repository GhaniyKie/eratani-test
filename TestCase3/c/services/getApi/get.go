package getapi

import (
	"eratani/TestCase3/c/storage/models"

	"github.com/sirupsen/logrus"
)

type ResultData interface {
	GetDataResponse(req models.RequestData) (resp []models.ResponseQuery, err error)
}

type Executor struct {
	logger     *logrus.Logger
	ResultData ResultData
}

func NewGetApi(logger *logrus.Logger, ResultData ResultData) *Executor {
	return &Executor{
		logger:     logger,
		ResultData: ResultData,
	}
}

func (e *Executor) LogicServices(req models.RequestData) (resp models.ResponseServices, err error) {
	logger := e.logger.WithFields(logrus.Fields{
		"Func Name": "LogicServices",
	})

	// Get data from layer storage mysql
	resultQuery, err := e.ResultData.GetDataResponse(req)
	if err != nil {
		logger.Errorf(`Error get result data query in services`)
		return resp, err
	}

	// If logic, create here

	resp.Data = resultQuery
	return
}
