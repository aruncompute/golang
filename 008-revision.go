package main
import "fmt"

type hotdog string

func main(){
	var x int;
	var y string
	var z bool
	var x1 = []int{1,3};
	x1=append(x1,42)

	
	var x2 hotdog
	x2="hotdog"


	x=42;
	fmt.Println(x);

	y = "james bond";
	fmt.Println(y);

	z=true
	fmt.Println(z);
	fmt.Println(x,y,z);
	fmt.Sprintf("X=%v, Y=%v, Z=%v",x,y,z);

	
	fmt.Println(x1);
	fmt.Println(x1[0]);	
	fmt.Println(x2);	

	
}