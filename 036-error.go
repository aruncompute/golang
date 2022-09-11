package main
import (
	"fmt"	
	"log"
	"encoding/json"
)

 
type person struct{
	Fname string
	Lname string
	Address string
	Age	int
}

func main(){		
	 p1 := person{
		Fname: "James",
		Lname:  "Bond",
		Age:   32,
		Address: 	"#238, Palm city Sector",
	 }
	 p2 := person{
		Fname: "Ansh",
		Lname: "Gupta",
		Address: "#238, Palm city, Sector 127",
		Age: 12,		
	 }
	 people := []person{p1, p2}

	  
	 fmt.Println(people);	 	 
	 log.Fatalln(people);
	 //log.Panicln(people);

	 fmt.Println("=================Error===========")
	 bs, err := json.Marshal(`{"name":"John", "age":30, "car":ok}`)
	 if err != nil {
		 log.Println(err)
	 }
	 log.Fatalln(string(bs))
}

 