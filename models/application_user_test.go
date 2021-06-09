package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDevicesPairs(t *testing.T) {

	var devices map[string]bool
	devices = make(map[string]bool)
	user1pair := ApplicationUser{
		UserId:     "1",
		NumDesktop: 1,
		NumLaptops: 0,
		DeviceIds:  devices,
	}

	user1pairLaptop := ApplicationUser{
		UserId:     "2",
		NumDesktop: 0,
		NumLaptops: 1,
		DeviceIds:  devices,
	}

	user2pairs := ApplicationUser{
		UserId:     "3",
		NumDesktop: 1,
		NumLaptops: 2,
		DeviceIds:  devices,
	}

	assert.Equal(t, 1, user1pair.CountDevicesPairs())
	assert.Equal(t, 1, user1pairLaptop.CountDevicesPairs())
	assert.Equal(t, 2, user2pairs.CountDevicesPairs())

}
