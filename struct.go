package hoproxy

// Config TODO あとで書く
type Config struct {
	Server    *Server     `yaml:"server"`
	Exchanges []*Exchange `yaml:"exchange"`
}

// Server TODO あとで書く
type Server struct {
	Port int `yaml:"port"`
}

// Exchange TODO あとで書く
type Exchange struct {
	Method  string `yaml:"method"`
	Path    string `yaml:"path"`
	CallURI string `yaml:"call_uri"`
	CallCli string `yaml:"call_cli"`
}

// Call TODO あとで書く
type Call struct {
	URI string
	Cli string
}

// ExchangeTree TODO あとで書く
type ExchangeTree struct {
	GET    map[string]*Call
	POST   map[string]*Call
	PUT    map[string]*Call
	DELETE map[string]*Call
	OTHER  map[string]*Call
}

// ExchangeResponse TODO あとで書く
type ExchangeResponse struct {
	StatusCode int
	Body       string
}

// ExchangeCli TODO あとで書く
type ExchangeCli struct {
	exitStatus int
	Body       string
	Error      error
}
