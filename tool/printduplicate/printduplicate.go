// Package -
package printduplicate

import (
	"errors"
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Doc explaining the tool.
const Doc = "Tool to enforce usage of our own internal file-writing utils instead of os.MkdirAll or ioutil.WriteFile"

var errUnsafePackage = errors.New(
	"os and ioutil dir and file writing functions are not permissions-safe, use fileutil",
)

// Analyzer runs static analysis.
var Analyzer = &analysis.Analyzer{
	Name:     "printduplicate",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("analyzer is not type *inspector.Inspector")
	}

	nodeFilter := []ast.Node{
		(*ast.MapType)(nil),
	}

	inspect.Preorder(nodeFilter, func(node ast.Node) {
		switch expr := node.(type) {
		case *ast.MapType:
			pass.Reportf(
				node.Pos(),
				fmt.Sprintf(
					"Found map: K=%s V=%s",
					expr.Key,
					expr.Value,
				),
			)
		}
	})

	return nil, nil
}
