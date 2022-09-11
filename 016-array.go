package main
import "fmt"

func main(){
	x1 := 5
	var x2 int = 5

	var y1 [5]int; 
	var y2 =[7]int{1,2,3,4,5}
	var y3 =[2][2]int{
		{1,2},
		{3,4},
	}
	
	fmt.Println(x1);
	fmt.Println(x2);
	fmt.Println(y1);
	fmt.Println(y2);
	fmt.Println(y3);
	fmt.Println(y2[1:3])

	fmt.Println("==Range==");
	for j, v :=range y2{
		fmt.Println(j,v);
	}
	
}