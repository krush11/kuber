package config

import (
	"net/http"
)

var BaseURL string = "https://www.nseindia.com"

/*
You cannot directly call all data from NSE API. Server checks for specific cookies in each req before processing the request.
To fetch the request, send the cookie data which are sent as a response from `/get-quotes/equity`.
Cookies are copied from response of `/get-quotes/equity` to further request which will be used to some verification on NSE's servers.
This function does that initial setup and sends a modified *http.Request that can be used directly for futher data fetching.

Cookies needed for further requests:
1. nsit
2. ak_bmsc
3. bm_sv
4. nseappid

How to fetch the cookies in Postman:
1. Call the `/get-quotes/equity` API twice.
2. Copy the cookies from the response of the second call.
3. Paste the cookies in the request header of the `/api/historical/cm/equity` API.
4. Send the request and you will get the data.
*/
func ReqConfig() *http.Request {
	req, _ := http.NewRequest("GET", BaseURL, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "www.nseindia.com")
	req.Header.Add("Referer", "https://www.nseindia.com/get-quotes/equity")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	for _, cookie := range res.Cookies() {
		req.AddCookie(cookie)
	}

	// TODO: Remove the need to call this API twice. This is just a temporary fix.
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	for _, cookie := range res.Cookies() {
		req.AddCookie(cookie)
	}

	// remove duplicate cookies from request.
	// This is the case when the cookie is set twice.
	// Needs to be removed once a permanent fix for fetching cookies is found.
	cookies := req.Cookies()
	for i := 0; i < len(cookies); i++ {
		for j := i + 1; j < len(cookies); j++ {
			if cookies[i].Name == cookies[j].Name {
				cookies = append(cookies[:j], cookies[j+1:]...)
				j--
			}
		}
	}
	req.Header.Del("Cookie")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	return req
}
