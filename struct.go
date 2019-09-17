package hoproxy

type Config struct {
	Server    *Server     `yaml:"server"`
	Exchanges []*Exchange `yaml:"exchange"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Exchange struct {
	Method  string `yaml:"method"`
	Path    string `yaml:"path"`
	CallUri string `yaml:"call_uri"`
	CallCli string `yaml:"call_cli"`
}

type Call struct {
	Uri string
	Cli string
}

type ExchangeTree struct {
	GET    map[string]*Call
	POST   map[string]*Call
	PUT    map[string]*Call
	DELETE map[string]*Call
	OTHER  map[string]*Call
}

type ExchangeResponse struct {
	StatusCode int
	Body       string
}

type ExchangeCli struct {
	exitStatus int
	Body       string
	Error      error
}
