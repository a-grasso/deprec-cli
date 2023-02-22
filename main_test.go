package main

import _ "net/http/pprof"

import (
	"os"
	"testing"
)

func Test(t *testing.T) {

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{
		"", "--runMode", "parallel", "test.sbom.json",
	}

	main()
}
