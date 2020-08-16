package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}
	t.Log("slice:[2:1]", slice[:3], slice)
	slice = append(slice, "90")
	t.Log("slice", slice)
	slice2 := make([]string, 10)
	copy(slice2, slice)
	t.Log("slice2", slice2)
}

func TestMakeAndNew(t *testing.T) {
	m1 := new([]int)
	m1 = &[]int{1,2,3}
	m2 := make([]int, 1)
	t.Log(m1, m2[0], *m1)
}

func TestSlice1(t *testing.T) {
	var number1 [] string
	number2 := []string{"a", "b", "c"}
	number3 := make([] int, 2, 2)
	number4 := append(number3, []int{3,4,5}...)
	number5 := [5]int{5,4,3,2,1}
	number6 := []string{"A", "B", "C"}
	number7 := copy(number2, number6)
	number8 := []int{3,2,4,56,2,16,74,21}
	number9 := number8
	number10 := append([]int{}, number8...)
	number9[0] = 0
	number10[0] = 1
	t.Log("number1", number1)
	t.Log("number2", number2, len(number1), cap(number1))
	t.Log("number3", number3, len(number3), cap(number3))
	t.Log("number4", number4)
	t.Log("number5", number5)
	t.Log("number7", number7, number2)
	t.Log("number9 and number8", number9, number8)
	t.Log("number10 and number8", number10, number8)

}

func TestSlice2(t *testing.T) {
	number1 := make([] int, 1, 2)
	fmt.Printf("%p\n", number1)
	fmt.Printf("%p\n", &number1[0])
	fmt.Printf("%p\n", &number1)
}