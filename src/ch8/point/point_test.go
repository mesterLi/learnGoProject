package point

import "testing"

type Book struct {
	Id string
	Name string
}
func TestPoint(t *testing.T) {
	book1 := Book{
		Id: "321",
		Name: "蒲公英女孩",
	}
	var book2 *Book = &book1
	var book3 Book = book1
	book2.Id = "123"
	book3.Id = "456"
	t.Log(*book2, book1)
	t.Log(book3, book1)
}

func TestPoint1(t *testing.T) {
	p := 123
	p2 := p
	p2 = 789
	var p1 *int // 生命p是int指针类型
	t.Log(p, &p)
	t.Log("空指针", &p1)
	t.Log(p2, &p2)
	// 内存地址不同
	p1 = &p
	t.Log(*p1, p1)
	t.Log(p, &p)
	*p1 = 456
	t.Log(*p1, p1)
	t.Log(p, &p)
}