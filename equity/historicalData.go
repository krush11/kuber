package equity

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/krush11/kuber/config"
)

type DailyDataType struct {
	ID                  string  `json:"_id"`
	Symbol              string  `json:"CH_SYMBOL"`
	Series              string  `json:"CH_SERIES"`
	MarkeyType          string  `json:"CH_MARKET_TYPE"`
	High                float32 `json:"CH_TRADE_HIGH_PRICE"`
	Low                 float32 `json:"CH_TRADE_LOW_PRICE"`
	Open                float32 `json:"CH_OPENING_PRICE"`
	Close               float32 `json:"CH_CLOSING_PRICE"`
	LTP                 float32 `json:"CH_LAST_TRADED_PRICE"`
	PrevClose           float32 `json:"CH_PREVIOUS_CLS_PRICE"`
	TotalTradedQuantity int32   `json:"CH_TOT_TRADED_QTY"`
	TotalTradedValue    float64 `json:"CH_TOT_TRADED_VAL"`
	High52W             float32 `json:"CH_52WEEK_HIGH_PRICE"`
	Low52W              float32 `json:"CH_52WEEK_LOW_PRICE"`
	TotalTrades         int32   `json:"CH_TOTAL_TRADES"`
	ISIN                string  `json:"CH_ISIN"`
	ChTimestamp         string  `json:"CH_TIMESTAMP"`
	Timestamp           string  `json:"TIMESTAMP"`
	CreatedAt           string  `json:"createdAt"`
	UpdatedAt           string  `json:"updatedAt"`
	V                   int8    `json:"__v"`
	VWAP                float32 `json:"VWAP"`
	MTimestamp          string  `json:"mTIMESTAMP"`
}

type SecurityDataType struct {
	Meta interface{} `json:"meta"`
	Data []DailyDataType `json:"data"`
}

/*
	Historical range bounded data fetch
	*	@param {string} symbol 	This is the symbol of security which you want to retrieve
	*	@param {string} from 		Starting date from which data needs to be fetched.
	*	@param {string} to 			Final data  till which data needs to be fetched
	*	@param {string} series 	Array of series of data required, separated by comma.

	Data returned in decreasing order of timestamp, i.e. data[0] will have latest date
*/
func FetchHistoricalData(symbol string, from string, to string, series string) (SecurityDataType, error) {
	req := config.ReqConfig()
	
	query := req.URL.Query()
	query.Add("symbol", symbol)
	query.Add("from", from)
	query.Add("to", to)
	query.Add("series", "[\""+series+"\"]")
	req.URL.RawQuery = query.Encode()
	req.URL.Path = "/api/historical/cm/equity"
	
	client := &http.Client{Timeout: 40 * time.Second}
	res, err := client.Do(req)
	var securityData SecurityDataType
	if err != nil {
		return securityData, err
	}
	defer res.Body.Close()

	// https://stackoverflow.com/questions/75248772/get-request-returns-data-in-thunder-client-postman-but-gives-blank-data-in-golan
	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
	default:
		reader = res.Body
	}
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&securityData)
	if err != nil {
		return securityData, err
	}

	// security := securityData.Data[0]
	// date, err := time.Parse(time.RFC3339, security.CreatedAt)
	// if err != nil {
	// 	log.Println(err)
	// }
	return securityData, nil
}
