package getapi

import (
	"eratani/TestCase3/c/storage/models"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type getApi struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewGetApi(logger *logrus.Logger, db *gorm.DB) *getApi {
	return &getApi{
		logger: logger,
		db:     db,
	}
}

func (g *getApi) GetDataResponse(req models.RequestData) (resp []models.ResponseQuery, err error) {
	logger := g.logger.WithFields(logrus.Fields{
		"Func Name": "GetDataResponse",
	})

	querySelect := `
		SELECT
			u.country, a.total_buy
		FROM
			data_belanja AS a
		INNER JOIN
			data_user AS u
		ON
			a.id_user = u.id
		GROUP BY
			u.country, a.total_buy
		ORDER BY
			a.total_buy DESC
	`

	queryConditions := []string{}

	if req.Id != 0 {
		queryConditions = append(queryConditions, fmt.Sprintf(` WHERE id = %d`, req.Id))
	}

	if len(queryConditions) > 0 {
		querySelect += strings.Join(queryConditions, " AND ")
	}

	querySelect += `
		ORDER BY
			id ASC
	`

	// Get data and mapping in struct
	dtResponse := []models.ResponseQuery{}
	if err := g.db.Raw(querySelect).Scan(&dtResponse).Error; err != nil {
		logger.Errorf(`Error query mysql`)
		return resp, errors.New("error query get data")
	}

	// Result struct mapping for response
	resp = dtResponse
	return resp, nil
}
