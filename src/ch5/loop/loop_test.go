package loop

import "testing"

func TestLoop(t *testing.T) {
	for a := 1; a <= 5; a++ {
		t.Log("a=", a)
	}
}

func TestLoop2(t *testing.T) {
	a := 1
	for a < 5 {
		if a == 2 {
			a++
			continue
		}
		t.Log("a=", a)
		a++
	}
}

func TestLoop3(t *testing.T) {
	a := 1
	for {
		a++
		t.Log("a=", a)
		if a >= 10 {
			break
		}
	}
}