package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var p Product // memory allocated and attributes are initialized
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 10
	*/

	// var p Product = Product{100, "Pen", 10}
	// var p Product = Product{Id: 100, Name: "Pen", Cost: 10}
	var p Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// create a copy
	/*
		p2 := p
		p2.Cost = 20
		fmt.Printf("p.Cost = %v, p2.Cost = %v\n", p.Cost, p2.Cost)
	*/

	// create a reference
	p2 := &p
	p2.Cost = 20
	fmt.Printf("p.Cost = %v, p2.Cost = %v\n", p.Cost, p2.Cost)

	fmt.Println("Before applying 10% discount:")
	fmt.Println(Format( /* ? */ ))
	ApplyDiscount( /* ? */ ) // apply 10% discount on the proeduct (p)
	fmt.Println("After applying 10% discount:")
	fmt.Println(Format( /* ? */ ))

}

func Format() string {
	// return a formatted string of the given product
	// ex : Id = 100, Name = "Pen", Cost = 10.00
}

func ApplyDiscount( /* ? */ ) { // no return values
	// update the cost of the given product by applying the given discount
}
