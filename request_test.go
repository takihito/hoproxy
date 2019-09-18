package hoproxy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var n int = 1000

var mockHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	age := r.Form.Get("age")
	i, _ := strconv.Atoi(age)
	i += n
	fmt.Fprintf(w, "%d", i)
})

func TestCallExchangeURI(t *testing.T) {

	ts := httptest.NewServer(mockHandler)
	defer ts.Close()

	getReq, _ := http.NewRequest("GET", "http://proxy-test-server/api/foo?age=99", nil)
	age := 99
	//	v := url.Values{}
	//	v.Set("age", strconv.Itoa(age))
	//	//	getReq.Form = v
	//	//req.Form.Encode()
	//	getReq.URL.RawQuery = v.Encode()

	exchangeResponse, _ := callExchangeURI(getReq, ts.URL)
	if exchangeResponse.StatusCode != 200 {
		t.Error("inalid GET res.StatusCode")
	}

	if exchangeResponse.Body != strconv.Itoa(age+n) {
		t.Error("inalid GET res.Body ")
	}

	age = 100
	v := url.Values{}
	v.Set("age", strconv.Itoa(age))
	postReq, _ := http.NewRequest("POST", "http://proxy-test-server/api/bar", strings.NewReader(v.Encode()))
	exchangeResponse, _ = callExchangeURI(postReq, ts.URL)
	if exchangeResponse.StatusCode != 200 {
		t.Error("inalid POST res.StatusCode")
	}
	if exchangeResponse.Body != strconv.Itoa(age+n) {
		t.Error("inalid POST res.Body")
	}

}
