package helpers

import (
	"log"

	"github.com/anebula/flex_go/models"
	"github.com/anebula/flex_go/utils"
)

/*
Given a csv record(row), either add the User with the information from the record
to the userIdMap, or if user already exists in the map with same Id, then update
user's desktop count, laptop count and list of devices.
Returns nothing, but modifies the input map of userIds to models.ApplicationUser.
*/
func CreateOrUpdateUser(record models.CsvRecord, userIdMap map[string]*models.ApplicationUser) {

	if _, ok := userIdMap[record.UserId]; ok {
		// if user is already on the userIdMap, check if this is a new device.
		if _, found := (*userIdMap[record.UserId]).DeviceIds[record.ComputerId]; !found {
			// update device counter accordingly for new device
			if record.ComputerType == utils.Desktop {
				(*userIdMap[record.UserId]).NumDesktop += 1
			} else if record.ComputerType == utils.Laptop {
				(*userIdMap[record.UserId]).NumLaptops += 1
			} else {
				log.Fatal("Unknown device" + record.ComputerType)
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
			log.Fatal("Unknown device" + record.ComputerType)
		}
		userIdMap[record.UserId] = &user
	}

}
