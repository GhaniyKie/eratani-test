package datacountry

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type RespDataCountry struct {
	Country    string `json:"country"`
	TotalSpend int    `json:"total_spend"`
}

type getCountry struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewGetCountry(logger *logrus.Logger, db *gorm.DB) *getCountry {
	return &getCountry{
		logger: logger,
		db:     db,
	}
}

func (g *getCountry) ResponseData() (resp []RespDataCountry, err error) {
	logger := g.logger.WithFields(logrus.Fields{
		"Func Name": "ResponseData",
	})

	query := `
		SELECT 
			b.country, 
			SUM(b.country) AS total_spend
		FROM 
			table_name AS b
		GROUP BY
			country 
		ORDER BY 
			total_spend DESC
	`

	// Siapkan struct nya
	sliceDataCountry := []RespDataCountry{}
	// Ambil hasil query dan mapping kedalam struct
	if err := g.db.Raw(query).Scan(&sliceDataCountry).Error; err != nil {
		logger.Errorf(`Error query`)
		return resp, errors.New("error query, check again")
	}

	// Mapping hasil result query kedalam response
	resp = sliceDataCountry
	return resp, nil
}
