package hoproxy

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateExcangeTree(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	data := `---
exchange:
  - method: get
    path: /api/get
    call_cli: /usr/local/bin/get.sh
  - method: put
    path: /api/put
    call_uri: http://example.local/api/put
    call_cli: /usr/local/bin/put.sh
  - method: post
    path: /api/post1
    call_cli: /usr/local/bin/post1.sh
  - method: post
    path: /api/post2
    call_uri: http://example.local/api/post2
    call_cli: /usr/local/bin/post2.sh
  - method: delete
    path: /api/delete
    call_cli: /usr/local/bin/delete.sh
  - method: head
    path: /api/head
    call_uri: http://example.local/api/head
`

	if err := ioutil.WriteFile(f.Name(), []byte(data), 0644); err != nil {
		t.Error(err)
	}

	cfg := NewConfig(f.Name())

	tree := CreateExchangeTree(cfg)
	if tree.GET["/api/get"].Cli != "/usr/local/bin/get.sh" {
		t.Errorf("not tree.GET./api/get %s", tree.GET["/api/get"].Cli)
	}
	if tree.GET["/api/get"].Uri != "" {
		t.Errorf("not tree.GET./api/get Uri %s", tree.GET["/api/get"].Uri)
	}

	if tree.PUT["/api/put"].Cli != "/usr/local/bin/put.sh" {
		t.Errorf("not tree.PUT./api/put %s", tree.PUT["/api/put"].Cli)
	}
	if tree.PUT["/api/put"].Uri != "http://example.local/api/put" {
		t.Errorf("not tree.PUT./api/put Uri %s", tree.PUT["/api/put"].Uri)
	}

	if tree.POST["/api/post1"].Cli != "/usr/local/bin/post1.sh" {
		t.Errorf("not tree.POST./api/post1 Cli %s", tree.POST["/api/post2"].Cli)
	}
	if tree.POST["/api/post1"].Uri != "" {
		t.Errorf("not tree.POST./api/post1 Uri %s", tree.POST["/api/post1"].Uri)
	}

	if tree.POST["/api/post2"].Cli != "/usr/local/bin/post2.sh" {
		t.Errorf("not tree.POST./api/post2 Cli %s", tree.POST["/api/post2"].Cli)
	}
	if tree.POST["/api/post2"].Uri != "http://example.local/api/post2" {
		t.Errorf("not tree.POST./api/post2 Uri %s", tree.POST["/api/post2"].Uri)
	}

	if tree.DELETE["/api/delete"].Cli != "/usr/local/bin/delete.sh" {
		t.Errorf("not tree.GET./api/delete %s", tree.DELETE["/api/delete"].Cli)
	}
	if tree.DELETE["/api/delete"].Uri != "" {
		t.Errorf("not tree.GET./api/delete Uri %s", tree.DELETE["/api/delete"].Uri)
	}

	if tree.OTHER["/api/head"].Cli != "" {
		t.Errorf("not tree.OTHER./api/head %s", tree.OTHER["/api/head"].Cli)
	}
	if tree.OTHER["/api/head"].Uri != "http://example.local/api/head" {
		t.Errorf("not tree.OTHER./api/head Uri %s", tree.OTHER["/api/head"].Uri)
	}

}
