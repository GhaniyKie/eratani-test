package getapi

import (
	"eratani/TestCase3/c/storage/models"
	"eratani/util/logger"
	"eratani/util/sqlmock"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_GetDataResponse(t *testing.T) {
	logger := logger.New()
	db, mock := sqlmock.NewMock()
	dbGorm, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatal(err)
	}

	request := models.RequestData{
		Id: 1,
	}

	querySelect := `
		SELECT
			*
		FROM
			table_name AS a
	`

	queryConditions := []string{}

	if request.Id != 0 {
		queryConditions = append(queryConditions, fmt.Sprintf(` WHERE id = %d`, request.Id))
	}

	if len(queryConditions) > 0 {
		querySelect += strings.Join(queryConditions, " AND ")
	}

	querySelect += `
		ORDER BY
			id ASC
	`

	rows := mock.NewRows([]string{"id", "country", "credit_card_type", "credit_card", "first_name", "last_name"}).
		AddRow(1, "indonesia", "mastercard", 5100133512561250, "testfirstname", "testlastname")

	mock.ExpectQuery(regexp.QuoteMeta(querySelect)).WillReturnRows(rows)
	getData := NewGetApi(logger, dbGorm)
	data, err := getData.GetDataResponse(request)
	assert.NotNil(t, data)
	assert.NoError(t, err)
}
