package smock

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go/types"
	"testing"
)

func TestOnFuncGeneration(t *testing.T) {
	expected := `
package mocks

import "github.com/stretchr/testify/mock"






	type CCall struct {
		*mock.Call
	}

	func (c *CCall) Return(
) *CCall {
		return &CCall{Call: c.Call.Return(
)}
	}
	
	func (_m *MyType) OnCMatch(matchers ...interface{}) *CCall {
		c := _m.On("C", matchers)
		return &CCall{Call: c}
	}
	
	func (_m *MyType) OnC(
		i int,
) *CCall {
		c := _m.On("C", 
		i,
)
		return &CCall{Call: c}
	}
`
	iface, err := loadType(".", "MyType")
	assert.NoError(t, err)
	var buf bytes.Buffer
	assert.NoError(t, onFuncTemplate.Execute(&buf, FuncSet{
		TypeName: iface.(*types.Named).Obj().Name(),
		Functions: map[FuncName]Function{
			"C": {
				Args: []Arg{
					{Name: "i", Type: iface.Underlying().(*types.Interface).Method(0).Type().(*types.Signature).Params().At(0).Type().String()},
				},
				ReturnArgs: []Arg{},
			},
		},
	}))

	assert.Equal(t, expected, buf.String())
}

func TestNewFuncSet(t *testing.T) {
	expected := FuncSet{
		TypeName: "MyType",
		Functions: map[FuncName]Function{
			"C": {
				Args: []Arg{
					{Name: "intvalue", Type: "int"},
				},
				ReturnArgs: []Arg{},
			},
		},
	}

	iface, err := loadType(".", "MyType")
	assert.NoError(t, err)
	fs := NewFuncSet(iface.(*types.Named))
	assert.NotNil(t, fs)
	assert.Equal(t, expected, fs)
}

func TestEndToEnd(t *testing.T) {
	iface, err := loadType(".", "DifferentCases")
	assert.NoError(t, err)
	fs := NewFuncSet(iface.(*types.Named))
	assert.NotNil(t, fs)

	expected := `
package mocks

import "github.com/stretchr/testify/mock"






	type ComplexArgsCall struct {
		*mock.Call
	}

	func (c *ComplexArgsCall) Return(
) *ComplexArgsCall {
		return &ComplexArgsCall{Call: c.Call.Return(
)}
	}
	
	func (_m *DifferentCases) OnComplexArgsMatch(matchers ...interface{}) *ComplexArgsCall {
		c := _m.On("ComplexArgs", matchers)
		return &ComplexArgsCall{Call: c}
	}
	
	func (_m *DifferentCases) OnComplexArgs(
		t smock.MyType,
) *ComplexArgsCall {
		c := _m.On("ComplexArgs", 
		t,
)
		return &ComplexArgsCall{Call: c}
	}
	type NoArgNamesCall struct {
		*mock.Call
	}

	func (c *NoArgNamesCall) Return(
) *NoArgNamesCall {
		return &NoArgNamesCall{Call: c.Call.Return(
)}
	}
	
	func (_m *DifferentCases) OnNoArgNamesMatch(matchers ...interface{}) *NoArgNamesCall {
		c := _m.On("NoArgNames", matchers)
		return &NoArgNamesCall{Call: c}
	}
	
	func (_m *DifferentCases) OnNoArgNames(
		stringvalue string,
) *NoArgNamesCall {
		c := _m.On("NoArgNames", 
		stringvalue,
)
		return &NoArgNamesCall{Call: c}
	}
	type ReturnNamesCall struct {
		*mock.Call
	}

	func (c *ReturnNamesCall) Return(
		err error,
) *ReturnNamesCall {
		return &ReturnNamesCall{Call: c.Call.Return(
		err,
)}
	}
	
	func (_m *DifferentCases) OnReturnNamesMatch(matchers ...interface{}) *ReturnNamesCall {
		c := _m.On("ReturnNames", matchers)
		return &ReturnNamesCall{Call: c}
	}
	
	func (_m *DifferentCases) OnReturnNames(
) *ReturnNamesCall {
		c := _m.On("ReturnNames", 
)
		return &ReturnNamesCall{Call: c}
	}
	type ReturnsCall struct {
		*mock.Call
	}

	func (c *ReturnsCall) Return(
		stringvalue string,
		errorvalue error,
) *ReturnsCall {
		return &ReturnsCall{Call: c.Call.Return(
		stringvalue,
		errorvalue,
)}
	}
	
	func (_m *DifferentCases) OnReturnsMatch(matchers ...interface{}) *ReturnsCall {
		c := _m.On("Returns", matchers)
		return &ReturnsCall{Call: c}
	}
	
	func (_m *DifferentCases) OnReturns(
		x string,
		i int,
) *ReturnsCall {
		c := _m.On("Returns", 
		x,
		i,
)
		return &ReturnsCall{Call: c}
	}
`
	var buf bytes.Buffer
	assert.NoError(t, onFuncTemplate.Execute(&buf, fs))
	assert.Equal(t, expected, buf.String())
}
