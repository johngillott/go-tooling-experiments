package main

import (
	"github.com/johngillott/go-tooling-experiments/tool/printduplicate"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(printduplicate.Analyzer)
}
