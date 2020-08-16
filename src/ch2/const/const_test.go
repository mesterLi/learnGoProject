package _const

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	WriteAble
	DelAble
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, WriteAble, DelAble)
}
