package handlers

import (
	"testing"

	"github.com/anebula/flex_go/models"
	"github.com/anebula/flex_go/utils"
	"github.com/stretchr/testify/assert"
)

/*
Test for reading csv concurrently
*/
func TestReadCsvConcurrent(t *testing.T) {

	test_csv := "../data/simple.csv"
	test_app_id := "999"

	expected_csv_rows := []models.CsvRecord{
		{ComputerId: "1", UserId: "U1", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "2", UserId: "U1", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "3", UserId: "U1", ApplicationId: "999", ComputerType: utils.Laptop, Comment: "Exported from System A"},
		{ComputerId: "10", UserId: "U2", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "10", UserId: "U2", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "11", UserId: "U2", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
		{ComputerId: "111", UserId: "U3", ApplicationId: "999", ComputerType: utils.Laptop, Comment: "Exported from System A"},
		{ComputerId: "112", UserId: "U3", ApplicationId: "999", ComputerType: utils.Desktop, Comment: "Exported from System A"},
	}

	csv_rows := ReadCsvConcurrent(test_csv, test_app_id)
	assert.Equal(t, len(expected_csv_rows), len(csv_rows))
}
