package postapi

import (
	"encoding/json"
	"eratani/TestCase3/d/storage/models"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/renderer"
)

type ResultServices interface {
	LogicServices(req models.RequestPost) (resp models.ResponseData, err error)
}

type Handler struct {
	logger         *logrus.Logger
	render         *renderer.Render
	ResultServices ResultServices
}

func NewHandler(logger *logrus.Logger, render *renderer.Render, ResultServices ResultServices) *Handler {
	return &Handler{
		logger:         logger,
		render:         render,
		ResultServices: ResultServices,
	}
}

func (h *Handler) HandlerPost(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithFields(logrus.Fields{
		"Func Name": "HandlerGet",
	})

	reqBody := models.RequestPost{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		logger.Errorf(`Error decode body`)
		h.render.JSON(w, http.StatusInternalServerError, &models.ResponseData{
			Message: "Error decode body",
			Data:    http.StatusInternalServerError,
		})
		return
	}

	// Validate conditions
	if reqBody.CreditCard != 0 && reqBody.CreditCardType == "" {
		logger.Errorf(`Error, column cant nil`)
		h.render.JSON(w, http.StatusBadRequest, &models.ResponseData{
			Message: "Error, column cant nil",
			Data:    http.StatusBadRequest,
		})
		return
	}

	// Get data from layer services
	resultServices, err := h.ResultServices.LogicServices(reqBody)
	if err != nil {
		logger.Errorf(`Error: %v`, err)
		h.render.JSON(w, http.StatusInternalServerError, &models.ResponseData{
			Message: "Error something",
			Data:    http.StatusInternalServerError,
		})
		return
	}

	h.render.JSON(w, http.StatusOK, resultServices)
}
