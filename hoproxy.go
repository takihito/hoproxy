package hoproxy

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

var (
	exchanges []*Exchange
	tree      *ExchangeTree
)

func Run(cfg *Config) {
	exchanges = cfg.Exchanges

	tree = CreateExchangeTree(cfg)

	http.HandleFunc("/", hellohandlerProxy)

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
}

func hellohandlerProxy(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context() // TODO 外部接続するのでつかう
	if tree.GET[r.URL.Path] != nil {
		callExchange(w, r, tree.GET[r.URL.Path])
	}
	if tree.POST[r.URL.Path] != nil {
		callExchange(w, r, tree.POST[r.URL.Path])
	}
	if tree.PUT[r.URL.Path] != nil {
		callExchange(w, r, tree.PUT[r.URL.Path])
	}
	if tree.DELETE[r.URL.Path] != nil {
		callExchange(w, r, tree.DELETE[r.URL.Path])
	}

	// TODO あとで差し替える
	fmt.Fprint(w, quote.Hello())
}

func callExchange(w http.ResponseWriter, r *http.Request, call *Call) {
	if call.Uri != "" {
		exchangeResponse, err := callExchangeUri(r, call.Uri) // あとでerrハンドリング
		if err != nil {
			w.WriteHeader(500)
			fmt.Println(err)
		}
		fmt.Fprint(w, exchangeResponse.Body)

	}
	if call.Cli != "" {
		exchangeCliResponse, err := callExchangeCli(r, call.Cli) // あとでerrハンドリング
		if err != nil {
			w.WriteHeader(500)
			fmt.Println(err)
		}
		fmt.Fprint(w, exchangeCliResponse.Body)
	}
}
