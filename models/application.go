package models

type Application struct {
	ApplicationId string
	Users         map[string]*ApplicationUser
}
