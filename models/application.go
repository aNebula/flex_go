package models

/*
Data model to represent an applicaiton
*/
type Application struct {
	ApplicationId string
	Users         map[string]*ApplicationUser // maps userIds to User model, for all users using this application
}
