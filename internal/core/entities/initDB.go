package entities

type InitDBFunc interface {
	Init() (err error)
}
