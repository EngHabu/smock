package mocks

import "github.com/stretchr/testify/mock"

type GetAbcCall struct {
	*mock.Call
}

func (c *GetAbcCall) Return(
	errorvalue error,
	intvalue int,
) *GetAbcCall {
	return &GetAbcCall{Call: c.Call.Return(
		errorvalue,
		intvalue,
	)}
}

func (_m *Simple) OnGetAbcMatch(matchers ...interface{}) *GetAbcCall {
	c := _m.On("GetAbc", matchers)
	return &GetAbcCall{Call: c}
}

func (_m *Simple) OnGetAbc() *GetAbcCall {
	c := _m.On("GetAbc",
	)
	return &GetAbcCall{Call: c}
}

type SetAbcCall struct {
	*mock.Call
}

func (c *SetAbcCall) Return() *SetAbcCall {
	return &SetAbcCall{Call: c.Call.Return()}
}

func (_m *Simple) OnSetAbcMatch(matchers ...interface{}) *SetAbcCall {
	c := _m.On("SetAbc", matchers)
	return &SetAbcCall{Call: c}
}

func (_m *Simple) OnSetAbc(
	intvalue int,
	stringvalue string,
) *SetAbcCall {
	c := _m.On("SetAbc",
		intvalue,
		stringvalue,
	)
	return &SetAbcCall{Call: c}
}
