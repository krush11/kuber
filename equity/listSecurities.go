package equity

import (
	"encoding/csv"
	"net/http"
	"strconv"
)

type SecurityType struct {
	Symbol      string
	CompanyName string
	Series      string
	ListingDate string
	PaidUpValue int8
	MarketLot   int8
	ISIN        string
	FaceValue   int8
}

func FetchSecurities() ([]SecurityType, error) {
	res, err := http.Get("https://archives.nseindia.com/content/equities/EQUITY_L.csv")
	var securitiesList []SecurityType
	if err != nil {
		return securitiesList, err
	}
	defer res.Body.Close()

	reader := csv.NewReader(res.Body)
	data, err := reader.ReadAll()
	if err != nil {
		return securitiesList, err
	}

	for i := range data {
		if i != 0 {
			var security SecurityType
			security.Symbol = data[i][0]
			security.CompanyName = data[i][1]
			security.Series = data[i][2]
			security.ListingDate = data[i][3]
			PaidUpValue, _ := strconv.Atoi(data[i][4])
			security.PaidUpValue = int8(PaidUpValue)
			MarketLot, _ := strconv.Atoi(data[i][4])
			security.MarketLot = int8(MarketLot)
			security.ISIN = data[i][6]
			FaceValue, _ := strconv.Atoi(data[i][4])
			security.FaceValue = int8(FaceValue)

			securitiesList = append(securitiesList, security)
		}
	}

	return securitiesList, err
}
