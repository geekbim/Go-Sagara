package response_test

import (
	"os"
	"sagara/domain/entity"
	"sagara/internal/delivery/response"
	"sagara/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOfficeResponse(t *testing.T) {
	os.Setenv("SERVER_PORT", "localhost:8080")
	os.Setenv("DB_TIMEZONE", "Asia/Jakarta")
	office := testdata.NewOffice()
	offices := []*entity.Office{office}

	res := response.MapOfficeListDomainToResponse(offices, int32(len(offices)))

	assert.Equal(t, offices[0].Id.String(), res.Offices[0].Id)
	assert.Equal(t, offices[0].Name, res.Offices[0].Name)
}

func TestOfficeResponse(t *testing.T) {
	os.Setenv("SERVER_PORT", "localhost:8080")
	os.Setenv("DB_TIMEZONE", "Asia/Jakarta")
	office := testdata.NewOffice()

	res := response.MapOfficeDomainToResponse(office)

	assert.Equal(t, office.Id.String(), res.Id)
	assert.Equal(t, office.Name, res.Name)
}
