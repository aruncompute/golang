package main
import "fmt"

func main(){
	var x1 = []int{4, 5, 7, 8, 42}
	x1 = append(x1,77, 88, 99, 1014)
	
	fmt.Println(x1);
}