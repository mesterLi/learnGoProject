package condition

import "testing"

func TestCondition(t *testing.T) {
	if a:=1; a == 1 {
		t.Log("a==1")
	} else {
		t.Log("a!=1")
	}
}
