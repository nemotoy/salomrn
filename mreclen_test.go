package mreclen_test

import (
	"testing"

	"github.com/nemotoy/mreclen"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, mreclen.Analyzer, "a")
}
