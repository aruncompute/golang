package main
import "fmt"

type Person struct{
	fname string
	lname string
	favColor []string
}
func main(){	
	fmt.Println("===============Exercise 1===============");
	x1 := Person{
		fname: "Arun ",
		lname: "Kumar",
		favColor: []string{"Red","Blue","Pink"},
	}
	fmt.Println(x1);

}