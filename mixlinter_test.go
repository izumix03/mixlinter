package mixlinter_test

import (
	"github.com/izumix03/mixlinter"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, mixlinter.Analyzer, "a")
}

