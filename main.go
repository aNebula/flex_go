package main

import (
	"fmt"

	"github.com/anebula/flex_go/handlers"
	"github.com/anebula/flex_go/helpers"
)

func main() {
	const csvFilename string = "data/sample-small.csv"
	const appIdFilter string = "999"

	csv_rows := handlers.ReadCsvConcurrent(csvFilename, appIdFilter)
	appCount := helpers.CountAppSubs(appIdFilter, csv_rows)

	fmt.Println(appCount)
}
