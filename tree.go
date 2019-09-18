package hoproxy

import (
	"strings"
)

// CreateExchangeTree TODO あとで書く
func CreateExchangeTree(cfg *Config) *ExchangeTree {
	tree := &ExchangeTree{
		GET:    map[string]*Call{},
		PUT:    map[string]*Call{},
		POST:   map[string]*Call{},
		DELETE: map[string]*Call{},
		OTHER:  map[string]*Call{},
	}

	for i := range cfg.Exchanges {
		call := &Call{
			URI: cfg.Exchanges[i].CallURI,
			Cli: cfg.Exchanges[i].CallCli,
		}
		method := strings.ToUpper(cfg.Exchanges[i].Method)
		switch {
		case method == "GET":
			tree.GET[cfg.Exchanges[i].Path] = call
		case method == "PUT":
			tree.PUT[cfg.Exchanges[i].Path] = call
		case method == "POST":
			tree.POST[cfg.Exchanges[i].Path] = call
		case method == "DELETE":
			tree.DELETE[cfg.Exchanges[i].Path] = call
		default:
			tree.OTHER[cfg.Exchanges[i].Path] = call
		}

	}

	return tree
}
