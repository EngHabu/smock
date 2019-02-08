
package mocks

import "github.com/stretchr/testify/mock"






	type GetAbcCall struct {
		*mock.Call
	}

	func (c *GetAbcCall) Return(
		intvalue int,
		errorvalue error,
) *GetAbcCall {
		return &GetAbcCall{Call: c.Call.Return(
		intvalue,
		errorvalue,
)}
	}
	
	func (_m *Simple) OnGetAbcMatch(matchers ...interface{}) *GetAbcCall {
		c := _m.On("GetAbc", matchers)
		return &GetAbcCall{Call: c}
	}
	
	func (_m *Simple) OnGetAbc(
) *GetAbcCall {
		c := _m.On("GetAbc", 
)
		return &GetAbcCall{Call: c}
	}
	type SetAbcCall struct {
		*mock.Call
	}

	func (c *SetAbcCall) Return(
) *SetAbcCall {
		return &SetAbcCall{Call: c.Call.Return(
)}
	}
	
	func (_m *Simple) OnSetAbcMatch(matchers ...interface{}) *SetAbcCall {
		c := _m.On("SetAbc", matchers)
		return &SetAbcCall{Call: c}
	}
	
	func (_m *Simple) OnSetAbc(
		stringvalue string,
		intvalue int,
) *SetAbcCall {
		c := _m.On("SetAbc", 
		stringvalue,
		intvalue,
)
		return &SetAbcCall{Call: c}
	}
