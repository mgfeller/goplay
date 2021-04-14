package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

func isErrorCodeName(name string) bool {
	matched, _ := regexp.MatchString("^Err[A-Z]", name)
	return matched
}

func walk(path string) error {
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

func inspect(path string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		spec, ok := n.(*ast.ValueSpec)
		if ok {
			for _, id := range spec.Names {
				if isErrorCodeName(id.Name) {
					value0 := id.Obj.Decl.(*ast.ValueSpec).Values[0]
					switch value := value0.(type) {
					case *ast.BasicLit:
						fmt.Printf("code value: %s\n", value.Value)
						value.Value = "\"hello\""
						//codeValue = strings.Trim(value.Value, "\"")
					}
				}
			}
		}
		return true
	})

	buf := new(bytes.Buffer)
	err = format.Node(buf, fset, file)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else if path[len(path)-8:] != "_test.go" {
		d, f := filepath.Split(path)
		ioutil.WriteFile(filepath.Join(d, "result_"+f), buf.Bytes(), 0644)
		//ioutil.WriteFile(path, buf.Bytes(), 0644)
	}
}

func main() {
	inspect("./test_files/one_var.go")
}
