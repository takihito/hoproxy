package hoproxy

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	data := `---
server:
  port: 8088
exchange:
  - method: post
    path: /api/post
    call_uri: http://example.local/api/post
    call_cli: /usr/local/bin/post.sh
  - method: post
    path: /api/delete
    call_cli: /usr/local/bin/delete.sh
`

	if err := ioutil.WriteFile(f.Name(), []byte(data), 0644); err != nil {
		t.Error(err)
	}

	cfg := NewConfig(f.Name())

	if cfg.Server.Port != 8088 {
		t.Error("not ok cfg.Server.Port ")
	}

	if cfg.Exchanges[0].Method != "post" {
		t.Error("not ok cfg.Exchanges[0].Method")
	}
	if cfg.Exchanges[0].Path != "/api/post" {
		t.Error("not ok cfg.Exchanges[0].Path")
	}
	if cfg.Exchanges[0].CallUri != "http://example.local/api/post" {
		t.Error("not ok cfg.Exchanges[0].CallUri")
	}
	if cfg.Exchanges[0].CallCli != "/usr/local/bin/post.sh" {
		t.Error("not ok cfg.Exchanges[0].CallCli")
	}
	if cfg.Exchanges[1].Path != "/api/delete" {
		t.Error("not ok cfg.Exchanges[0].Path")
	}
	if cfg.Exchanges[1].CallCli != "/usr/local/bin/delete.sh" {
		t.Error("not ok cfg.Exchanges[1].CallCli")
	}

}
