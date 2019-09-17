package hoproxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func callExchangeUri(req *http.Request, uri string) (*ExchangeResponse, error) {
	exReq, err := http.NewRequest(req.Method, uri, req.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if req.Method == "GET" {
		exReq.URL.RawQuery = req.URL.RawQuery
	}

	exReq.ContentLength = req.ContentLength

	if req.Method == "POST" {
		exReq.Form = req.Form
		exReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := new(http.Client)
	exRes, err := client.Do(exReq)
	defer exRes.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	exchangeResponse := &ExchangeResponse{}
	body, error := ioutil.ReadAll(exRes.Body)
	if error != nil {
		log.Fatal(error)
	}
	exchangeResponse.Body = string(body)
	exchangeResponse.StatusCode = exRes.StatusCode

	return exchangeResponse, nil
}
