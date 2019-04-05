package main

import (
	"github.com/izumix03/mixlinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(mixlinter.Analyzer)
}
