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
	Args       map[VarName]SimpleTypeName
	ReturnArgs map[VarName]SimpleTypeName
}

type FuncSet struct {
	Functions map[FuncName]Function
	TypeName  SimpleTypeName
}

func NewFunction(f *types.Signature) Function {
	args := make(map[VarName]SimpleTypeName, f.Params().Len())
	for i := 0; i < f.Params().Len(); i++ {
		v := f.Params().At(i)
		varName := v.Name()
		if len(varName) == 0 {
			varName = fmt.Sprintf("%vvalue", v.Type().String())
		}

		args[varName] = v.Type().String()
	}

	returnArgs := make(map[VarName]SimpleTypeName, f.Results().Len())
	for i := 0; i < f.Results().Len(); i++ {
		v := f.Results().At(i)
		varName := v.Name()
		if len(varName) == 0 {
			varName = fmt.Sprintf("%vvalue", v.Type().String())
		}

		returnArgs[varName] = v.Type().String()
	}

	return Function{Args: args, ReturnArgs: returnArgs}
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
{{define "argsDef"}}
	{{- range $varName, $varType := .}}
		{{$varName}} {{$varType}},
	{{- end}}
{{end}}

{{define "args"}}
	{{- range $varName, $varType := .}}
		{{$varName}},
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

type DifferentCases interface {
	Returns(x string, i int) (string, error)
	NoArgNames(string)
	ReturnNames() (err error)
	ComplexArgs(t MyType)
}
