package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/johngillott/go-tooling-experiments/tool/printdistinct"
)

func main() {
	singlechecker.Main(printdistinct.Analyzer)
}
