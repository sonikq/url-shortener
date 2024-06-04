package main

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestRun(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, osExitAnalyzer, "./...")
}
