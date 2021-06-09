package models

import "github.com/anebula/flex_go/utils"

/*
Data model to represent an user of applicaiton
*/
type ApplicationUser struct {
	UserId string

	NumLaptops int // number of laptops that has applciation installed

	NumDesktop int // number of desktops that has application installed

	DeviceIds map[string]bool // a map of the device ids of the user with the application installed
}

func (appUser ApplicationUser) CountDevicesPairs() int {
	return utils.Min(appUser.NumDesktop, appUser.NumLaptops) + utils.Abs(appUser.NumLaptops-appUser.NumDesktop)
}
