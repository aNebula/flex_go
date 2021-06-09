package models

import "github.com/anebula/flex_go/utils"

type ApplicationUser struct {
	UserId string

	NumLaptops int

	NumDesktop int

	DeviceIds map[string]bool
}

func (appUser ApplicationUser) CountUnpairedDevices() int {
	return min(appUser.NumDesktop, appUser.NumLaptops) + utils.Abs(appUser.NumLaptops-appUser.NumDesktop)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
