package getapi

import (
	"eratani/TestCase3/c/storage/models"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/renderer"
)

const (
	msgInternalServerError = "Something error"
	msgBadRequest          = "Error, Bad Request"
	msgSuccess             = "Success"
)

type ResultServices interface {
	LogicServices(req models.RequestData) (resp models.ResponseServices, err error)
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

func (h *Handler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithFields(logrus.Fields{
		"Func Name": "HandlerGet",
	})

	// Create param input and convert to int
	param := r.URL.Query()
	Id, err := strconv.Atoi(param.Get("id"))
	if err != nil {
		logger.Errorf(`Params input ID must int.`)
		h.render.JSON(w, http.StatusBadRequest, &models.ResponseHandler{
			Message: msgBadRequest,
			Items:   http.StatusBadRequest,
		})
		return
	}

	request := models.RequestData{
		Id: Id,
	}

	// Get data from layer services
	resultServices, err := h.ResultServices.LogicServices(request)
	if err != nil {
		logger.Errorf(`Error get data result services in handler.`)
		h.render.JSON(w, http.StatusInternalServerError, &models.ResponseHandler{
			Message: msgInternalServerError,
			Items:   http.StatusInternalServerError,
		})
		return
	}

	h.render.JSON(w, http.StatusOK, &models.ResponseHandler{
		Message: msgSuccess,
		Items:   resultServices,
	})
}
