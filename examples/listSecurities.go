package examples

import (
	"log"

	"github.com/krush11/kuber/equity"
)

func listEquitySecurities() {
	securitiesList, _ := equity.FetchSecurities()
	log.Println(securitiesList)
}