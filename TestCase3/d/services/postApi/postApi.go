package postapi

import (
	"eratani/TestCase3/d/storage/models"
	"errors"

	"github.com/sirupsen/logrus"
)

type ResultQueryInsert interface {
	PostData(req models.RequestPost) (resp models.ResponsePost, err error)
}

type Executor struct {
	logger            *logrus.Logger
	ResultQueryInsert ResultQueryInsert
}

func NewExecutor(logger *logrus.Logger, ResultQueryInsert ResultQueryInsert) *Executor {
	return &Executor{
		logger:            logger,
		ResultQueryInsert: ResultQueryInsert,
	}
}

func (e *Executor) LogicServices(req models.RequestPost) (resp models.ResponseData, err error) {
	logger := e.logger.WithFields(logrus.Fields{
		"Layer": "Logic Services",
	})

	// Get data from layer storage
	resultQuery, err := e.ResultQueryInsert.PostData(req)
	if err != nil {
		logger.Errorf(`Error get result query`)
		resp.Message = "Error"
		resp.Data = nil
		return resp, errors.New("error get result query insert data")
	}

	resp.Message = "Success"
	resp.Data = resultQuery
	return resp, nil
}
