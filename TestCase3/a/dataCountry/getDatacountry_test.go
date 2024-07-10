package datacountry

import (
	"eratani/util/logger"
	"eratani/util/sqlmock"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_ResponseData(t *testing.T) {
	logger := logger.New()
	db, mock := sqlmock.NewMock()
	dbGorm, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatal(err)
	}

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

	// Buat mock response
	rows := mock.NewRows([]string{"country", "total_spend"}).AddRow("indonesia", 10)

	// Check mock query dan hasilnya
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// Penuhi constructor dari function
	getCountry := NewGetCountry(logger, dbGorm)
	// Check return
	data, err := getCountry.ResponseData()
	// Check hasil
	assert.NotNil(t, data)
	assert.NoError(t, err)
}
