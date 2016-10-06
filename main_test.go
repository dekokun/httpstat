package main

import (
	"strings"
	"testing"
)

func Test引数がないとエラーになる(t *testing.T) {
	args := strings.Split("httpstat", " ")
	if status := Run(args); status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}

func Test存在しないステータスコードだとエラーになる(t *testing.T) {
	args := strings.Split("httpstat 8888", " ")
	if status := Run(args); status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}
}

func Test404はOK(t *testing.T) {
	args := strings.Split("httpstat 404", " ")
	if status := Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
