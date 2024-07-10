package postapi

import (
	"eratani/TestCase3/d/storage/models"
	"eratani/util/logger"
	"eratani/util/sqlmock"
	"errors"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_PostData(t *testing.T) {
	logger := logger.New()
	db, mock := sqlmock.NewMock()
	dbGorm, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatal(err)
	}

	request := models.RequestPost{
		Country:        "Indonesia",
		CreditCardType: "mastercard",
		CreditCard:     12345,
		FirstName:      "firstname",
		LastName:       "lastname",
	}

	query := `
		INSERT INTO
		table_name
			(country, credit_card_type, credit_card, first_name, last_name)
		VALUES
			(?, ?, ?, ?, ?)
	`

	rows := mock.NewRows([]string{"country", "credit_card_type",
		"credis_card", "first_name", "last_name"}).AddRow("Indonesia", "mastercard",
		12345, "firstname", "lastname")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	postData := NewPost(logger, dbGorm)
	data, err := postData.PostData(request)
	if err != nil {
		errors.New("error insert query, check again.")
	}
	assert.NotNil(t, data)
}
