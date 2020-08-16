package bubbleSort

import "testing"

func TestBubbleSort(t *testing.T) {
	list := []int{2,1,4,5,7,9,12,10}
	var a []int
	t.Log(a)
	if a == nil {
		t.Log("a === nil")
	}
	for i := 0; i < len(list) - 1; i++ {
		for j := i+1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
		//t.Log(list[i])
	}
	t.Log(list)
}
