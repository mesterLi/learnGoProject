package struct_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Student struct {
	Gradge int `json:"gradge"`
	Subject string `json:"subject"`
	*People `json:"people"`
}

type People struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (p *People) sayHello() {
	fmt.Println(p.Name, "say hello")
}

func TestStruct(t *testing.T)  {
	var student Student
	p := People{
		Name: "alice",
		Age: 18,
	}
	p.sayHello()
	student.Gradge = 1
	student.Subject = "123"
	student.People = &People{
		Name: "xiaoming",
		Age: 18,
	}
	data, _ := json.Marshal(student)
	fmt.Print(student, student.People, "\n")
	fmt.Print("data ", string(data), "\n")
}