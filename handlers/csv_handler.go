package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/anebula/flex_go/models"
)

/*
Reading csv concurrently saves about 4s on 1Gb file. Spawning more goroutine for filtering and
instantiating models.CSVRecords does not improve performance. This is because the processing task is
small and much quicker than disk I/O.
*/
func ReadCsvConcurrent(filename string, appIdFilter string) []models.CsvRecord {

	csv_file, err := os.Open(filename)
	defer csv_file.Close()
	if err != nil {
		log.Fatal("Error opening csv file. " + err.Error())
	}

	log.Println("Parsing csv now...")

	records := make(chan []string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		parser := csv.NewReader(csv_file)
		defer close(records)
		for {
			record, err := parser.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Error reading row from the csv.")
			}

			records <- record
		}
	}()

	var (
		mu   sync.Mutex
		rows []models.CsvRecord
	)

	rows = make([]models.CsvRecord, 0)

	go func() {
		for aRecord := range records {
			if aRecord[2] == appIdFilter {
				newRow := models.CsvRecord{
					ComputerId:    aRecord[0],
					UserId:        aRecord[1],
					ApplicationId: aRecord[2],
					ComputerType:  strings.ToLower(strings.TrimSpace(aRecord[3])),
					Comment:       aRecord[4],
				}
				mu.Lock()
				rows = append(rows, newRow)
				mu.Unlock()
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return rows
}
