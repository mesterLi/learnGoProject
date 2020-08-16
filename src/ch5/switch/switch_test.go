package _switch

import "testing"

func TestSwitch(t *testing.T) {
	i := 1
	for i < 10 {
		switch i {
			case 1,2,3:
				t.Log("123")
			case 4,5,6:
				t.Log("4,5,6")
		default:
			t.Log("default")
		}
		i++
	}
}

type Book struct {
	name,desc string
	id int
}
func TestRange(t *testing.T) {
	list := []Book{
		{name: "三体", id: 1, desc: "ascacasc"},
		{name: "夜的第七章", id: 2, desc: "ascacascasc"},
	}
	for _,v := range list {
		//t.Log("k=======", _)
		t.Log("v=======", v, "\n")
	}
}