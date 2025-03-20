package notes

import "log"

type Animal interface {
	Eat()
}
type Cat struct {
	Name string
}

func (c Cat) Eat() {
	log.Println(c.Name, "吃")
}

var L Animal

func Interface() {
	c := Cat{
		Name: "Tom",
	}
	Myfun(c)
	L.Eat()
}

func Myfun(a Animal) {
	L = a
}
