package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/anebula/flex_go/handlers"
	"github.com/anebula/flex_go/helpers"
)

// struct for reading config
type Config struct {
	Filename      string
	ApplicationId string
}

func main() {

	// read config
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal("Error decoding config toml. Please check valid toml is present at ./config.toml")
	}

	csv_rows := handlers.ReadCsvConcurrent(conf.Filename, conf.ApplicationId)
	log.Println("Parsing csv completed.")
	appCount := helpers.CountAppSubs(conf.ApplicationId, csv_rows)

	fmt.Println("You need to order minimum " + strconv.Itoa(appCount) + " copies of application with ID " + conf.ApplicationId)
}
