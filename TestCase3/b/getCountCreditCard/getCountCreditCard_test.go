package getcountcreditcard

import (
	"eratani/util/logger"
	"eratani/util/sqlmock"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_ResponseQuery(t *testing.T) {
	logger := logger.New()
	db, mock := sqlmock.NewMock()
	dbGorm, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatal(err)
	}

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

	// Buat mock response
	rows := mock.NewRows([]string{"credit_card_type", "total"}).AddRow("mastercard", 100)

	// Check mock query dan hasilnya
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	// Penuhi constructor dari function
	getCountry := NewGetCountData(logger, dbGorm)
	// Check return
	data, err := getCountry.ResponseQuery()
	// Check hasil
	assert.NotNil(t, data)
	assert.NoError(t, err)
}
