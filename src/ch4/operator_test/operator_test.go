package operator_test

import "testing"

func TestCompareArrar(t *testing.T) {
	a := [...]string{"A", "B", "C"}
	b := [...]string{"C", "B", "A"}
	c := [...]string{"A", "B", "C"}
	t.Log(a == b)
	t.Log(a == c)
	t.Log(b == c)
}
