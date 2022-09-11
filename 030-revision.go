package main
import (
	"fmt"	
)

 
type person struct{
	fname string
	lname string
	address string
}
func main(){		
	fmt.Println("=================Exercise 1=================")
	x :=100;
	fmt.Printf("Value of x is %v, Address of x %v \n", x,&x);

	fmt.Println("=================Exercise 2=================")
	p1 := person{
		fname: "Arun",
		lname: "Kumar",
		address: "#83, Azad Nagar Balong",
	}
	changeMe(&p1,"#238, Palm city")
	p1.details();
}

func changeMe(p *person, s string){
	p.address=s
}
func (p person) details(){
	fmt.Printf("Name: %v %v \n",p.fname, p.lname)
	fmt.Printf("Address: %v",p.address)
}