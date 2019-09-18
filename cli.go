package hoproxy

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// MaxOutputLength TODO あとで書く
var MaxOutputLength = 1000

// CommandTimeout TODO あとで書く
var CommandTimeout = 60 * time.Second

func callExchangeCli(req *http.Request, command string) (*ExchangeCli, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CommandTimeout)
	defer cancel()

	var commands []string
	var exchangeCli = &ExchangeCli{}
	if strings.Index(command, " ") != -1 {
		commands = []string{"sh", "-c", command}
	} else {
		commands = []string{command}
	}

	var input string
	if req.Body != nil {
		body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576)) // 1MiB
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		input = string(body)
	}

	//	q := u.Query()
	q := req.URL.Query()
	//	for k := range req.Form {
	for k := range q {
		commands = append(commands, k)
		commands = append(commands, q.Get(k))
	}

	cmd := exec.CommandContext(ctx, commands[0], commands[1:]...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return exchangeCli, err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		exchangeCli.Error = err
		fmt.Println(err)
	}

	exchangeCli.exitStatus = 0
	exchangeCli.Body = string(output)

	return exchangeCli, err
}
