package salomrn_test

import (
	"testing"

	"github.com/nemotoy/salomrn"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, salomrn.Analyzer, "a")
}
