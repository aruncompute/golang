package main
import (
	"fmt"
	"encoding/json"
	"log"
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

	 bs, err := json.Marshal(people)
	 if err != nil {
		 fmt.Println(err)
	 }
	 fmt.Println(string(bs))

	 /*//////////////////// Unmarshell //////////////*/
	decodeData	:= []person{};
	 err1 := json.Unmarshal([]byte(bs), &decodeData);
	 if err1 != nil {
		 fmt.Println(err1)
	 }
	 fmt.Println(people);
	 fmt.Println("=========Logs====================")
	 log.Println(people);
	 log.Fatalln(people);
	 log.Panicln(people);
}

 