package hoproxy

import (
	"net/http"
	"strings"
	"testing"
)

func TestCallExchangeCli(t *testing.T) {

	getReq, _ := http.NewRequest("GET", "http://proxy-test-server/api/foo?-e=print 2", nil)
	//	code := "print 1 + 1"
	//	v := url.Values{}
	//	v.Set("-e", code)
	//	getReq.Form = v
	exchangeCli, _ := callExchangeCli(getReq, "perl")
	if exchangeCli.Body != "2" {
		t.Error("inalid GET exchangeCli.Body")
	}
	if exchangeCli.Error != nil {
		t.Error("incalid GET exchangeCli.Error")
	}

	json := "{\"age\":100}"
	postReq, _ := http.NewRequest("POST", "http://proxy-test-server/api/bar", strings.NewReader(json))
	exchangeCli, _ = callExchangeCli(postReq, "jq -c \".age + 1\"")
	if exchangeCli.Body == "101" {
		t.Error("inalid exchangeCli.Body")
	}
	if exchangeCli.Error != nil {
		t.Error("incalid POST exchangeCli.Error")
	}

	getReqErr, _ := http.NewRequest("GET", "http://proxy-test-server/api/foo?-e=exit 128", nil)
	//	code = "exit 128"
	//	v := url.Values{}
	//	v.Set("-e", code)
	//	getReqErr.Form = v
	exchangeCli, _ = callExchangeCli(getReqErr, "perl")
	if exchangeCli.Body != "" {
		t.Error("inalid exchangeCli.Body")
	}
	if exchangeCli.Error.Error() != "exit status 128" {
		t.Error("exchangeCli.Error")
	}

}
