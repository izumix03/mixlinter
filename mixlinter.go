package mixlinter

import (
	"flag"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"reflect"
	"strings"
)

var includeTest bool

func init() {
	Analyzer.Flags.BoolVar(&includeTest, "include test file", false, "include test file or not")
}

var Analyzer = &analysis.Analyzer{
	Name: "mixlinter",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Flags: flag.FlagSet{
		Usage: nil,
	},
}

const Doc = "mixlinter is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}

	ins.Preorder(nodeFilter, func(n ast.Node) {
		fileDirList := strings.Split(pass.Fset.File(n.Pos()).Name(), "/")
		fileName := fileDirList[len(fileDirList) - 1]
		if strings.HasPrefix(fileName, "mock_") {
			return
		}
		if !includeTest && strings.HasSuffix(fileName, "_test") {
			return
		}
		switch n := n.(type) {
		case *ast.CompositeLit:
			var fields []string
			var setFields []string

			if reflect.ValueOf(n).IsNil() {
				return
			}
			if ident, ok := n.Type.(*ast.Ident); ok {
				if reflect.ValueOf(ident.Obj).IsNil() || reflect.ValueOf(ident.Obj.Decl).IsNil() {
					return
				}
				if ts, ok := ident.Obj.Decl.(*ast.TypeSpec); ok {
					if st, ok := ts.Type.(*ast.StructType); ok {
						for _, f := range st.Fields.List {
							fields = append(fields, f.Names[0].Name)
						}
					}
				}
			}
			for _, k := range n.Elts {
				if kv, ok := k.(*ast.KeyValueExpr); ok {
					if ident, ok := kv.Key.(*ast.Ident); ok {
						setFields = append(setFields, ident.Name)
					}
				}
			}
			for _, f := range fields {
				if !contain(f, setFields) {
					pass.Reportf(n.Pos(), "uninitialised field found: %s", f)
				}
			}
		}
	})

	return nil, nil
}

func contain(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}
