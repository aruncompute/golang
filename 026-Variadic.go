package main
import "fmt"

func main(){		

	sum := sumOfNNumber(10,101,10,10);	
	fmt.Println("Sum of all numbers: ",sum);
}


func sumOfNNumber(number ...int) (int){
	sum := 0;
	for _,val := range(number){
		sum = sum+val;
	}
	
	return sum
}