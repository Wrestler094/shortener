// Package noosexit содержит анализатор, запрещающий использование os.Exit в функции main.
package noosexit

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const (
	mainPkgName  = "main"
	mainFuncName = "main"
	osPkgName    = "os"
	exitFuncName = "Exit"
)

var Analyzer = &analysis.Analyzer{
	Name: "noosexit",
	Doc:  "запрещает вызов os.Exit в функции main пакета main",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if pass.Pkg.Name() != mainPkgName {
		return nil, nil
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if isOsExitCall(call) && isInsideMainFunc(file, call.Pos()) {
				reportOsExitUsage(pass, call)
			}

			return true
		})
	}

	return nil, nil
}

func isOsExitCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkg, ok := sel.X.(*ast.Ident)
	return ok && pkg.Name == osPkgName && sel.Sel.Name == exitFuncName
}

func reportOsExitUsage(pass *analysis.Pass, call *ast.CallExpr) {
	pos := pass.Fset.Position(call.Pos())
	pass.Reportf(call.Pos(), "запрещено использовать os.Exit в функции main (%s:%d)", pos.Filename, pos.Line)
}

func isInsideMainFunc(file *ast.File, pos token.Pos) bool {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == mainFuncName && fn.Body != nil {
			return fn.Body.Pos() <= pos && pos <= fn.Body.End()
		}
	}
	return false
}
