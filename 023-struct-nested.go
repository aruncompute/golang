package main
import "fmt"

type Person struct {
	name string	
	age   int
	address  string
}

type SpecialAgent struct{
	Person,
	ltk string
}

func main(){
	/*x1 := Person {
		name: "Arun kumar",
		age: 37,
		address: "#238, Palm city, Secotr 127, Mohali",
	}*/

	/*ag1 := SpecialAgent{
		Person: Person{name: "Arun kumar",age: 37,address: "#238, Palm city, Secotr 127, Mohali"},
		ltk: "true"
	}*/

	ag1 := SpecialAgent{Person: Person{name: "10", age: 10, address: "Hello world"}, ltk: "foo"}
	fmt.Println(ag1);

}