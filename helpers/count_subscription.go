package helpers

import "github.com/anebula/flex_go/models"

func CountAppSubs(appId string, csv_rows []models.CsvRecord) int {
	var userIdMap map[string]*models.ApplicationUser
	userIdMap = make(map[string]*models.ApplicationUser)

	currentApp := models.Application{
		ApplicationId: appId,
		Users:         userIdMap,
	}

	for _, row := range csv_rows {
		CreateOrUpdateUser(row, currentApp.Users)
	}

	appCount := 0

	for _, user := range currentApp.Users {
		user := *user
		appCount = appCount + user.CountDevicesPairs()
	}

	return appCount
}
