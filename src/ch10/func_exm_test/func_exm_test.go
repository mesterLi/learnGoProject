package func_exm_test

import (
	"fmt"
	"testing"
)

type Programer interface {
	sayHello() string
}
type goProgramer struct {

}
type javaProgramer struct {

}
func (gp *goProgramer) sayHello() string {
	return "goProgramer say hello"
}

func (jp *javaProgramer) sayHello() string {
	return "javaProgramer say hello"
}

func writeSayHello(p Programer) {
	fmt.Print("\nwriteSayHello", p.sayHello())
}
func TestFunc(t *testing.T) {
	gp := new(goProgramer)
	jp := new(javaProgramer)
	t.Log("\n", gp.sayHello())
	t.Log("\n", jp.sayHello())
	writeSayHello(gp)
	writeSayHello(jp)

}