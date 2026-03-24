package ibc

type Line interface {
	Defaults()
	ReadLine(input string) error
	Load(rc []string) error
	ToString() string
}
