package getcountcreditcard

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type RespDataCount struct {
	CreditCardType string `json:"credit_card_type"`
	Total          int    `json:"total"`
}

type GetCountData struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewGetCountData(logger *logrus.Logger, db *gorm.DB) *GetCountData {
	return &GetCountData{
		logger: logger,
		db:     db,
	}
}

func (g *GetCountData) ResponseQuery() (resp []RespDataCount, err error) {
	logger := g.logger.WithFields(logrus.Fields{
		"Func Name": "ResponseQuery",
	})

	query := `
		SELECT 
			credit_card_type , 
			COUNT(*) AS total 
		FROM 
			table_name  
		GROUP BY 
			credit_card_type  
		ORDER BY 
			total DESC 
		LIMIT 
			1;
	`

	sliceCountData := []RespDataCount{}
	if err := g.db.Raw(query).Scan(&sliceCountData).Error; err != nil {
		logger.Errorf(`Error query`)
		return resp, errors.New("error query, check again")
	}

	resp = sliceCountData
	return resp, nil
}
