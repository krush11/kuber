package examples

import (
	"log"
	"time"

	"github.com/krush11/kuber/equity"
)

func fetchEquityData() {
		symbol := "PAYTM"
	series := "EQ"
	from_date := time.Date(2023, 1, 20, 0, 0, 0, 0, time.Local).Format("02-01-2006")
	to_date := time.Date(2023, 1, 21, 0, 0, 0, 0, time.Local).Format("02-01-2006")

	historicalData, err := equity.FetchHistoricalData(symbol, from_date, to_date, series)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(historicalData)
	}
}