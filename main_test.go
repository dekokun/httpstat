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
