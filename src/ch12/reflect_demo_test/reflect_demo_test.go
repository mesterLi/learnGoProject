package reflect_demo_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Cat struct {
	Name string
	Age int
}

func (c *Cat) say() {
	fmt.Println(c.Name, "say 喵喵喵")
}

func TestReflectDemo(t *testing.T) {
	a := 1
	b := "1"
	c := []int{1}
	e := map[string]int{"a": 1}
	t.Log(typeOf(a))
	t.Log(typeOf(b))
	t.Log(typeOf(c))
	t.Log(typeOf(e))
	t.Log("================分割线==============")
	t.Log(valueOf(a))
	t.Log(valueOf(b))
	t.Log(valueOf(c))
	t.Log(valueOf(e))
}

func TestInterface(t *testing.T) {
	//c := Cat{
	//	Name: "果酱",
	//	Age: 2,
	//}
	cv := reflect.ValueOf(1)
	v := cv.Interface();
	fmt.Println(v)
	i := 1
	l := reflect.ValueOf(&i)
	fmt.Println(l)
	l.Elem().SetInt(10);
	fmt.Println(i)
}

func typeOf(d interface{}) interface{} {
	return reflect.TypeOf(d)
}

func valueOf(d interface{}) interface{} {
	return reflect.ValueOf(d)
}
