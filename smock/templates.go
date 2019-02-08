package smock

import (
	"fmt"
	"go/types"
	"text/template"
)

type VarName = string
type FuncName = string
type SimpleTypeName = string

type Function struct {
	Args       []Arg
	ReturnArgs []Arg
}

type Arg struct {
	Name VarName
	Type SimpleTypeName
}

type FuncSet struct {
	Functions map[FuncName]Function
	TypeName  SimpleTypeName
}

func NewFunction(f *types.Signature) Function {
	orderedArgs := make([]Arg, 0, f.Params().Len())
	for i := 0; i < f.Params().Len(); i++ {
		v := f.Params().At(i)
		varName := v.Name()
		if len(varName) == 0 {
			varName = fmt.Sprintf("%vvalue", afterLastDot(v.Type().String()))
		}

		orderedArgs = append(orderedArgs, Arg{Name: varName, Type: afterLastSlash(v.Type().String())})
	}

	orderedReturnArgs := make([]Arg, 0, f.Results().Len())
	for i := 0; i < f.Results().Len(); i++ {
		v := f.Results().At(i)
		varName := v.Name()
		if len(varName) == 0 {
			varName = fmt.Sprintf("%vvalue", afterLastDot(v.Type().String()))
		}

		orderedReturnArgs = append(orderedReturnArgs, Arg{Name: varName, Type: afterLastSlash(v.Type().String())})
	}

	return Function{
		Args:       orderedArgs,
		ReturnArgs: orderedReturnArgs,
	}
}

func NewFuncSet(st *types.Named) FuncSet {
	typeName := st.Obj().Name()

	funcs := make(map[FuncName]Function, st.NumMethods())
	iface := st.Underlying().(*types.Interface)
	for i := 0; i < iface.NumMethods(); i++ {
		m := iface.Method(i)
		funcs[m.Name()] = NewFunction(m.Type().(*types.Signature))
	}

	return FuncSet{
		TypeName:  typeName,
		Functions: funcs,
	}
}

var onFuncTemplate = template.Must(template.New("OnFuncs").Parse(`
package mocks

import "github.com/stretchr/testify/mock"

{{define "argsDef"}}
	{{- range $var := .}}
		{{$var.Name}} {{$var.Type}},
	{{- end}}
{{end}}

{{define "args"}}
	{{- range $var := .}}
		{{$var.Name}},
	{{- end}}
{{end}}

{{$TypeName := .TypeName}}

{{- range $name, $func := .Functions}}
	type {{$name}}Call struct {
		*mock.Call
	}

	func (c *{{$name}}Call) Return({{template "argsDef" $func.ReturnArgs }}) *{{$name}}Call {
		return &{{$name}}Call{Call: c.Call.Return({{template "args" $func.ReturnArgs }})}
	}
	
	func (_m *{{ $TypeName }}) On{{$name}}Match(matchers ...interface{}) *{{$name}}Call {
		c := _m.On("{{$name}}", matchers)
		return &{{$name}}Call{Call: c}
	}
	
	func (_m *{{ $TypeName }}) On{{$name}}({{template "argsDef" $func.Args }}) *{{$name}}Call {
		c := _m.On("{{$name}}", {{template "args" $func.Args }})
		return &{{$name}}Call{Call: c}
	}
{{- end}}
`))

type MyType interface {
	C(int)
}

////go:generate go run ..\cmd\smock\main.go DifferentCases

type DifferentCases interface {
	Returns(x string, i int) (string, error)
	NoArgNames(string)
	ReturnNames() (err error)
	ComplexArgs(t MyType)
}
