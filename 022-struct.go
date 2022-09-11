package main
import "fmt"

type Person struct {
	name string	
	age   int
	address  string
}

func main(){
	x1 := Person {
		name: "Arun kumar",
		age: 37,
		address: "#238, Palm city, Secotr 127, Mohali",
	}
	x2 := Person {
		name: "Ansh Gupta",
		age: 12,
		address: "#238, Palm city, Secotr 127, Mohali",
	}
	 
	fmt.Println("Name: ",x1.name);
	fmt.Println("Age: ",x1.age);
	fmt.Println("Address: ",x1.address);
	fmt.Println(x2);

}