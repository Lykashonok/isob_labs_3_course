package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/tools/refactor/rename"
)

func isGoFile(path string) bool {
	return filepath.Ext(path) == ".go"
}

func stringConstsToVar(path string) error {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	set := token.NewFileSet()
	file, err := parser.ParseFile(set, path, contents, 0)
	if err != nil {
		// If the file is invalid, we do nothing.
		return nil
	}

	ctv := &constToVar{}
	for _, decl := range file.Decls {
		ast.Walk(ctv, decl)
	}
	sort.Sort(ctv)

	var resBuf bytes.Buffer
	var lastIdx int
	for _, decl := range ctv.Decls {
		start := int(decl.Pos() - 1)
		end := int(decl.End() - 1)
		resBuf.Write(contents[lastIdx:start])
		declData := contents[start:end]
		varData := strings.Replace(string(declData), "const", "var", 1)
		resBuf.WriteString(varData)
		lastIdx = end
	}
	resBuf.Write(contents[lastIdx:])

	return ioutil.WriteFile(path, resBuf.Bytes(), 0755)
}

type constToVar struct {
	Decls []*ast.GenDecl
}

const hashedSymbolSize = 10

type NameHasher []byte

func (n NameHasher) Hash(token string) string {
	hashArray := sha256.Sum256(append(n, []byte(token)...))

	hexStr := strings.ToLower(hex.EncodeToString(hashArray[:hashedSymbolSize]))
	for i, x := range hexStr {
		if x >= '0' && x <= '9' {
			x = 'g' + (x - '0')
			hexStr = hexStr[:i] + string(x) + hexStr[i+1:]
		}
	}
	if strings.ToUpper(token[:1]) == token[:1] {
		hexStr = strings.ToUpper(hexStr[:1]) + hexStr[1:]
	}
	return hexStr
}

func (c *constToVar) Visit(n ast.Node) ast.Visitor {
	if decl, ok := n.(*ast.GenDecl); ok {
		if decl.Tok == token.CONST {
			if constOnlyHasStrings(decl) {
				c.Decls = append(c.Decls, decl)
			}
		}
	}
	return c
}

func (c *constToVar) Len() int {
	return len(c.Decls)
}

func (c *constToVar) Swap(i, j int) {
	c.Decls[i], c.Decls[j] = c.Decls[j], c.Decls[i]
}

func (c *constToVar) Less(i, j int) bool {
	return c.Decls[i].Pos() < c.Decls[j].Pos()
}

func constOnlyHasStrings(decl *ast.GenDecl) bool {
	for _, spec := range decl.Specs {
		if cs, ok := spec.(*ast.ValueSpec); ok {
			if !specIsString(cs) {
				return false
			}
		}
	}
	return true
}

func specIsString(v *ast.ValueSpec) bool {
	if v.Type != nil {
		s, ok := v.Type.(fmt.Stringer)
		if ok && s.String() == "string" {
			return true
		}
	}
	if len(v.Values) != 1 {
		return false
	}
	return exprIsString(v.Values[0])
}

func exprIsString(e ast.Expr) bool {
	switch e := e.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			return true
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			return exprIsString(e.X) || exprIsString(e.Y)
		}
		return false
	case *ast.ParenExpr:
		return exprIsString(e.X)
	}
	return false
}

func ObfuscateStrings(gopath string) error {
	return filepath.Walk(gopath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !isGoFile(path) {
			return nil
		}
		if err := stringConstsToVar(path); err != nil {
			return err
		}

		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		set := token.NewFileSet()
		file, err := parser.ParseFile(set, path, contents, 0)
		if err != nil {
			return nil
		}

		obfuscator := &stringObfuscator{Contents: contents}
		for _, decl := range file.Decls {
			ast.Walk(obfuscator, decl)
		}
		newCode, err := obfuscator.Obfuscate()
		if err != nil {
			return err
		}
		return ioutil.WriteFile(path, newCode, 0755)
	})
}

type stringObfuscator struct {
	Contents []byte
	Nodes    []*ast.BasicLit
}

func (s *stringObfuscator) Visit(n ast.Node) ast.Visitor {
	if lit, ok := n.(*ast.BasicLit); ok {
		if lit.Kind == token.STRING {
			s.Nodes = append(s.Nodes, lit)
		}
		return nil
	} else if decl, ok := n.(*ast.GenDecl); ok {
		if decl.Tok == token.CONST || decl.Tok == token.IMPORT {
			return nil
		}
	} else if _, ok := n.(*ast.StructType); ok {
		// Avoid messing with annotation strings.
		return nil
	}
	return s
}

func (s *stringObfuscator) Obfuscate() ([]byte, error) {
	sort.Sort(s)

	parsed := make([]string, s.Len())
	for i, n := range s.Nodes {
		var err error
		parsed[i], err = strconv.Unquote(n.Value)
		if err != nil {
			return nil, err
		}
	}

	var lastIndex int
	var result bytes.Buffer
	data := s.Contents
	for i, node := range s.Nodes {
		strVal := parsed[i]
		startIdx := node.Pos() - 1
		endIdx := node.End() - 1
		result.Write(data[lastIndex:startIdx])
		result.Write(obfuscatedStringCode(strVal))
		lastIndex = int(endIdx)
	}
	result.Write(data[lastIndex:])
	return result.Bytes(), nil
}

func (s *stringObfuscator) Len() int {
	return len(s.Nodes)
}

func (s *stringObfuscator) Swap(i, j int) {
	s.Nodes[i], s.Nodes[j] = s.Nodes[j], s.Nodes[i]
}

func (s *stringObfuscator) Less(i, j int) bool {
	return s.Nodes[i].Pos() < s.Nodes[j].Pos()
}

func obfuscatedStringCode(str string) []byte {
	var res bytes.Buffer
	res.WriteString("(func() string {\n")
	res.WriteString("mask := []byte(\"")
	mask := make([]byte, len(str))
	for i := range mask {
		mask[i] = byte(rand.Intn(256))
		res.WriteString(fmt.Sprintf("\\x%02x", mask[i]))
	}
	res.WriteString("\")\nmaskedStr := []byte(\"")
	for i, x := range []byte(str) {
		res.WriteString(fmt.Sprintf("\\x%02x", x^mask[i]))
	}
	res.WriteString("\")\nres := make([]byte, ")
	res.WriteString(strconv.Itoa(len(mask)))
	res.WriteString(`)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())`)
	return res.Bytes()
}

func ObfuscatePackageNames(gopath string, n NameHasher) error {
	ctx := build.Default
	ctx.GOPATH = gopath

	level := 1
	srcDir := filepath.Join(gopath, "src")

	doneChan := make(chan struct{})
	defer close(doneChan)

	for {
		resChan := make(chan string)
		go func() {
			scanLevel(srcDir, level, resChan, doneChan)
			close(resChan)
		}()
		var gotAny bool
		for dirPath := range resChan {
			gotAny = true
			isMain := isMainPackage(dirPath)
			encPath := encryptPackageName(dirPath, n)
			srcPkg, err := filepath.Rel(srcDir, dirPath)
			if err != nil {
				return err
			}
			dstPkg, err := filepath.Rel(srcDir, encPath)
			if err != nil {
				return err
			}
			if err := rename.Move(&ctx, srcPkg, dstPkg, ""); err != nil {
				return fmt.Errorf("package move: %s", err)
			}
			if isMain {
				if err := makeMainPackage(encPath); err != nil {
					return fmt.Errorf("make main package %s: %s", encPath, err)
				}
			}
		}
		if !gotAny {
			break
		}
		level++
	}

	return nil
}

func scanLevel(dir string, depth int, res chan<- string, done <-chan struct{}) {
	if depth == 0 {
		select {
		case res <- dir:
		case <-done:
			return
		}
		return
	}
	listing, _ := ioutil.ReadDir(dir)
	for _, item := range listing {
		if item.IsDir() {
			scanLevel(filepath.Join(dir, item.Name()), depth-1, res, done)
		}
		select {
		case <-done:
			return
		default:
		}
	}
}

func encryptPackageName(dir string, p NameHasher) string {
	subDir, base := filepath.Split(dir)
	return filepath.Join(subDir, p.Hash(base))
}

func isMainPackage(dir string) bool {
	listing, err := ioutil.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, item := range listing {
		if isGoFile(item.Name()) {
			path := filepath.Join(dir, item.Name())
			set := token.NewFileSet()
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				return false
			}
			file, err := parser.ParseFile(set, path, contents, 0)
			if err != nil {
				return false
			}
			fields := strings.Fields(string(contents[int(file.Package)-1:]))
			if len(fields) < 2 {
				return false
			}
			return fields[1] == "main"
		}
	}
	return false
}

func makeMainPackage(dir string) error {
	listing, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, item := range listing {
		if !isGoFile(item.Name()) {
			continue
		}
		path := filepath.Join(dir, item.Name())
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		set := token.NewFileSet()
		file, err := parser.ParseFile(set, path, contents, 0)
		if err != nil {
			return err
		}

		pkgNameIdx := int(file.Package) + len("package") - 1
		prePkg := contents[:pkgNameIdx]
		postPkg := string(contents[pkgNameIdx:])

		fields := strings.Fields(postPkg)
		if len(fields) < 1 {
			return errors.New("no fields after package keyword")
		}
		packageName := fields[0]

		var newData bytes.Buffer
		newData.Write(prePkg)
		newData.WriteString(strings.Replace(postPkg, packageName, "main", 1))

		if err := ioutil.WriteFile(path, newData.Bytes(), item.Mode()); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	files := []string{"ui.go", "orm.go", "app.go", "actions.go"}

	var n NameHasher
	buf := make([]byte, 32)
	rand.Read(buf)
	n = buf
	for _, file := range files {
		ObfuscateStrings(file)
		ObfuscatePackageNames(file, n)
	}
}
