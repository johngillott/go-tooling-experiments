package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {

	f, err := os.OpenFile("/Users/johngillott/go/src/github.com/johngillott/go-tooling-experiments/tool/printdistinct/example/main.go", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "demo", f, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(file, func(x ast.Node) bool {

		switch v := x.(type) {

		case (*ast.MapType):

			fmt.Printf("Found map %+v\n", v)
			return false

		default:
			// fmt.Printf("Found thing %+v\n", x)
			return true
		}
	})
}
