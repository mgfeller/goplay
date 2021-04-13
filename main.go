package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

func isErrorCodeName(name string) bool {
	matched, _ := regexp.MatchString("^Err[A-Z]", name)
	return matched
}

func transform(path string) error {
	logger := log.WithFields(log.Fields{"path": path})
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return err
	}
	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.FuncDecl:
			logger.WithFields(log.Fields{"decl": "func", "name": decl.Name.String()}).Debug("ast declaration")
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					logger.WithFields(log.Fields{"decl": "importspec", "name": strings.Trim(spec.Path.Value, "\"")}).Debug("ast declaration")
				case *ast.TypeSpec:
					logger.WithFields(log.Fields{"decl": "typespec", "name": spec.Name.String()}).Debug("ast declaration")
				case *ast.ValueSpec:
					for _, id := range spec.Names {
						if isErrorCodeName(id.Name) {
							value0 := id.Obj.Decl.(*ast.ValueSpec).Values[0]
							codeValue := ""
							switch value := value0.(type) {
							case *ast.BasicLit:
								codeValue = strings.Trim(value.Value, "\"")
								logger.WithFields(log.Fields{"name": id.Name, "value": codeValue}).Info("Err* variable detected with literal value.")
							case *ast.CallExpr:
								logger.WithFields(log.Fields{"name": id.Name}).Warn("Err* variable detected with call expression value.")
							}
						}
					}
				default:
					logger.Debug(fmt.Sprintf("unhandled token type: %s", decl.Tok))
				}
			}
		default:
			logger.Debug(fmt.Sprintf("unhandled declaration: %v at %v", decl, decl.Pos()))
		}
	}
	return nil
}

func main() {
	transform("./test_files/one_var.go")
}
