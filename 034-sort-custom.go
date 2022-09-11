package main
import (
	"fmt"
	"sort"
)

type person struct{
	name string
	age int
}

func main(){		
	p1 := person{
		name: "Arun",
		age:	34,
	}
	p2 := person{
		name: "Varun",
		age:	4,
	}
	p3 := person{
		name: "Tarun",
		age:	55,
	}
	p4 := person{
		name: "Sarun",
		age:	12,
	}
	p5 := person{
		name: "Marun",
		age:	28,
	}
	
	allP := []person{p1,p2,p3,p4,p5}
	sort.Slice(allP, func(i, j int) bool {
		return allP[i].age < allP[j].age
	  })	
	fmt.Println(allP)
}