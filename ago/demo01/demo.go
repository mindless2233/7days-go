package main

import "fmt"

type Person struct {
	name, addr, phone string
	age               uint
}

func (p Person) fp1() {
	p.age = 21

}
func (p *Person) fp2() {
	p.age = 21

}
func main() {
	p1 := new(Person)
	p1.name = "cjc"
	//fmt.Println(p1.name, p1.addr, p1.phone)

	p2 := &Person{
		name:  "c2",
		addr:  "历下",
		phone: "156xxxx5352",
		age:   24,
	}
	p3 := Person{
		name:  "c2",
		addr:  "历下",
		phone: "156xxxx5352",
		age:   24,
	}
	//	fmt.Printf("%+v", p2)

	fmt.Println(p1, p2, p3)
	fmt.Printf("%T , %T , %T\n", p1, p2, p3)
	fmt.Println("-----------------------------------")
	p3.fp2()
	fmt.Println(p3.age)
	p2.fp1()
	fmt.Println(p2.age)
}
