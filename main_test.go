package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test引数がないとエラーになる(t *testing.T) {
	args := strings.Split("httpstat", " ")
	if status := makeCli().Run(args); status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}

func Test存在しないステータスコードだとエラーになる(t *testing.T) {
	args := strings.Split("httpstat 8888", " ")
	if status := makeCli().Run(args); status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}
}

func Test404はOK(t *testing.T) {
	args := strings.Split("httpstat 404", " ")
	if status := makeCli().Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}

func makeCli() *CLI {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	return cli
}
