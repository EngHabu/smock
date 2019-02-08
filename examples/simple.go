package examples

//go:generate mockery -name Simple
//go:generate go run ../cmd/smock/main.go Simple

type Simple interface {
	SetAbc(string, int)
	GetAbc() (int, error)
}
