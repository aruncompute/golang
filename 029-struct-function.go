package main
import (
	"fmt"	
)

type person struct{
	fname string
	lname string
	age	int
	langs []string
}

func main(){		
	e1 := person{
		fname: "Arun",
		lname: "Kumar",
		age:	32,
		langs: []string{"Hindi","English","punjabi"},

	}
	e1.details();
	//fmt.Println(e1);

}

func (p person) details(){
	fmt.Println("Name=", p.fname, " " ,p.lname)
	fmt.Println("Age=",p.age)
	fmt.Println("Langues")
	for i,l :=range p.langs{		
		if i > 0{
			fmt.Print(",")
		}
		fmt.Print(l)
	}


	
}
