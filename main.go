package main

import (
	"fmt"
	"time"

	"github.com/anebula/flex_go/handlers"
	"github.com/anebula/flex_go/models"
	"github.com/anebula/flex_go/utils"
)

func createOrUpdateUser(record models.CsvRecord, userIdMap map[string]*models.ApplicationUser) {

	var user models.ApplicationUser

	if _, ok := userIdMap[record.UserId]; ok {
		// if user is already on the userIdMap, check if this is a new device.
		user = *userIdMap[record.UserId]

		if _, found := user.DeviceIds[record.ComputerId]; !found {
			// update device counter accordingly for new device
			if record.ComputerType == utils.Desktop {
				user.NumDesktop += 1
			} else if record.ComputerType == utils.Laptop {
				user.NumLaptops += 1
			} else {
				panic("Unknown device" + record.ComputerType)
			}

			user.DeviceIds[record.ComputerId] = true
		}

	} else {
		// if user is not in the userIdMap, create a new user
		user = models.ApplicationUser{
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
	const csvFilename string = "sample-large.csv"
	const appIdFilter string = "374"

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
	/*
		var (
			mu       sync.Mutex
			appCount int
		)
		appCount = 0

		for _, v := range currentApp.Users {
			user := *v
			appCount = appCount + user.CountUnpairedDevices()
		}

		duration := time.Since(start)


		users := gen(currentApp.Users)
		subscriptions := countUserSubscriptions(users)

		for n := range subscriptions {
			mu.Lock()
			appCount = appCount + n
			mu.Unlock()
		}
	*/

	appCount := 0

	for _, user := range currentApp.Users {
		user := *user
		appCount = appCount + user.CountUnpairedDevices()
	}

	duration := time.Since(start)
	fmt.Println(appCount)
	fmt.Println(duration)
}

/*
func gen(users map[string]*models.ApplicationUser) <-chan models.ApplicationUser {
	out := make(chan models.ApplicationUser)
	go func() {
		for _, n := range users {
			out <- *n
		}
		close(out)
	}()
	return out
}

func countUserSubscriptions(appUser <-chan models.ApplicationUser) chan int {
	out := make(chan int)
	go func() {
		for n := range appUser {
			out <- n.CountUnpairedDevices()
		}
		close(out)
	}()
	return out

}
*/
