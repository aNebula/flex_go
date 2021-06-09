package helpers

import (
	"testing"

	"github.com/anebula/flex_go/models"
	"github.com/anebula/flex_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateUser(t *testing.T) {
	csv_rows := []models.CsvRecord{
		{ComputerId: "1", UserId: "U1", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "2", UserId: "U1", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "3", UserId: "U1", ApplicationId: "999", ComputerType: utils.Laptop, Comment: "Exported from System A"},
		{ComputerId: "10", UserId: "U2", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "10", UserId: "U2", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "11", UserId: "U2", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "111", UserId: "U3", ApplicationId: "999", ComputerType: utils.Laptop, Comment: "Exported from System A"},
		{ComputerId: "112", UserId: "U3", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
	}

	count := CountAppSubs("999", csv_rows)

	assert.Equal(t, 5, count)
}
