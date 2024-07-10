package postapi

import (
	"eratani/TestCase3/d/storage/models"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Post struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewPost(logger *logrus.Logger, db *gorm.DB) *Post {
	return &Post{
		logger: logger,
		db:     db,
	}
}

func (p *Post) PostData(req models.RequestPost) (resp models.ResponsePost, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"Layer":     "Mysql",
		"Func Name": "PostData",
	})

	query := `
		INSERT INTO
		table_name
			(country, credit_card_type, credit_card, first_name, last_name)
		VALUES
			(?, ?, ?, ?, ?)
	`

	trx := p.db.Begin()
	defer trx.Rollback()

	if err := trx.Exec(query, req.Country, req.CreditCardType, req.CreditCard, req.FirstName, req.LastName).Error; err != nil {
		logger.Errorf(`Error insert query`)
		return resp, errors.New("error insert query, check again.")
	}

	// For get last insert ID
	var lastInserted struct {
		Id int `json:"id"`
	}
	if err := trx.Raw(`SELECT LAST_INSERT_ID() as id`).Find(&lastInserted).Error; err != nil {
		logger.Errorf("error get last inserted id insert data: %v", err)
		return resp, err
	}

	resp.Id = lastInserted.Id
	resp.Country = req.Country
	resp.CreditCardType = req.CreditCardType
	resp.CreditCard = req.CreditCard
	resp.FirstName = req.FirstName
	resp.LastName = req.LastName

	return resp, nil
}
