package method

import (
	"fmt"
	"testing"
	"unsafe"
)
type Employee struct {
	Name string
	Id string
	Age int
}
func (e Employee) String() string {
	//fmt.Println(e.Age)
	e.Age = 99
	fmt.Printf("address not pointer is %x", unsafe.Pointer(&e.Age))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

//func (e *Employee) String() string {
//	//fmt.Println(e.Age)
//	fmt.Printf("address is %x", unsafe.Pointer(&e.Age))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

func TestMethod(t *testing.T) {
	e := Employee{"小李", "acascas", 18}
	t.Log(e.String())
	fmt.Printf("address is %x \n", unsafe.Pointer(&e.Age))
	fmt.Println(e.Age)
}

