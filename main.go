package main

import (
	"fmt"
	"time"

	"github.com/anebula/flex_go/handlers"
	"github.com/anebula/flex_go/models"
	"github.com/anebula/flex_go/utils"
)

func createOrUpdateUser(record models.CsvRecord, userIdMap map[string]*models.ApplicationUser) {

	if _, ok := userIdMap[record.UserId]; ok {
		// if user is already on the userIdMap, check if this is a new device.
		if _, found := (*userIdMap[record.UserId]).DeviceIds[record.ComputerId]; !found {
			// update device counter accordingly for new device
			if record.ComputerType == utils.Desktop {
				(*userIdMap[record.UserId]).NumDesktop += 1
			} else if record.ComputerType == utils.Laptop {
				(*userIdMap[record.UserId]).NumLaptops += 1
			} else {
				panic("Unknown device" + record.ComputerType)
			}

			(*userIdMap[record.UserId]).DeviceIds[record.ComputerId] = true
		}

	} else {
		// if user is not in the userIdMap, create a new user
		user := models.ApplicationUser{
			UserId:     record.UserId,
			NumLaptops: 0,
			NumDesktop: 0,
			DeviceIds:  map[string]bool{record.ComputerId: true},
		}
		if record.ComputerType == utils.Desktop {
			user.NumDesktop += 1
		} else if record.ComputerType == utils.Laptop {
			user.NumLaptops += 1
		} else {
			panic("Unknown device" + record.ComputerType)
		}
		userIdMap[record.UserId] = &user
	}

}

func main() {
	const csvFilename string = "simple.csv"
	const appIdFilter string = "999"

	start := time.Now()

	csv_rows := handlers.ReadCsvConcurrent(csvFilename, appIdFilter)

	var userIdMap map[string]*models.ApplicationUser

	userIdMap = make(map[string]*models.ApplicationUser)
	currentApp := models.Application{
		ApplicationId: appIdFilter,
		Users:         userIdMap,
	}

	for _, row := range csv_rows {
		createOrUpdateUser(row, currentApp.Users)
	}

	appCount := 0

	for _, user := range currentApp.Users {
		user := *user
		appCount = appCount + user.CountUnpairedDevices()
	}

	duration := time.Since(start)
	fmt.Println(appCount)
	fmt.Println(duration)
}
