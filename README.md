# smock
Generate Mock's On&lt;method>() helper functions

This is a complimentary tool to https://github.com/vektra/mockery which generates Mock implementation
for interfaces.

e.g.
```go
package realcode

//go:generate mockery -name MyInterface
//go:generate smock MyInterface

type MyInterface interface {
	GetAbc(int) (string, error)
	SetValue(index int, val string) (err error)
}
```

This generates two files; mocks/MyInterface.go and mocks/MyInterface_on.go

And can be used as follows:
```go
package testcode

func TestMyOtherFunc(t *testing.T) {
	m := mocks.MyInterface{}
	m.OnGetAbc(5).Return("my_val", nil)
	m.OnSetValueWithMatchers(mock.Anything, mock.Anything).Return("", fmt.Errorf("Unexpected"))
	
	// run test code
}
```