package nofieldname_test

import (
	"testing"

	"github.com/moriuss/nofieldname"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, nofieldname.Analyzer, "a")
	analysistest.Run(t, testdata, nofieldname.Analyzer, "b")
}
