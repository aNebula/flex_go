package models

import "github.com/anebula/flex_go/utils"

type ApplicationUser struct {
	UserId string

	NumLaptops int

	NumDesktop int

	DeviceIds map[string]bool
}

func (appUser ApplicationUser) CountDevicesPairs() int {
	return utils.Min(appUser.NumDesktop, appUser.NumLaptops) + utils.Abs(appUser.NumLaptops-appUser.NumDesktop)
}
