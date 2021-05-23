# go-tooling-experiments

## package docs

- ["go/build"](https://pkg.go.dev/go/build)
- ["go/packages](https://pkg.go.dev/golang.org/x/tools/go/packages)
- ["go/parser"](https://pkg.go.dev/go/parser)
- ["golang.org/x/tools/go/analysis"](https://pkg.go.dev/golang.org/x/tools/go/analysis)

## blog posts

- [Using go/analysis to write a custom linter](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/)
- [Writing multi-package analysis tools for Go](https://eli.thegreenplace.net/2020/writing-multi-package-analysis-tools-for-go/)
- [Custom Static Analysis in Go (go/analysis)](https://rauljordan.com/2020/11/01/custom-static-analysis-in-go-part-1.html)

## working examples

- [Exercism's Go Analyzer](https://github.com/exercism/go-analyzer)
- [Staticcheck](https://github.com/dominikh/go-tools/blob/master/cmd/structlayout/main.go)

## glossary

### expression

Expressions only contain identifiers, literals and operators, where operators include arithmetic and boolean operators, the function call operator () the subscription operator [] and similar, and can be reduced to some kind of 'value', which can be any Python object. Examples:

```python
3 + 5
map(lambda x: x*x, range(10))
[a.x for a in some_iterable]
yield 7
```

### statements

Statements represent an action or command e.g print statements,
assignment statements. Expression is a combination of variables,
operations and values that yields a result value.

```python
print 42
if x: do_y()
return
a = 7
```
