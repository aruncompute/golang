package main
import "fmt"

func main(){

	map1 := map[int]int{
		1 : 1,
		2 : 1,
		3 : 1,
		4 : 1,
		5 : 1,
	}
	map1[6]=1;
	map2 := map[int]int{
		1 : 1,
		2 : 1,
		3 : 1,
		4 : 1,
		5 : 1,
	}

	delete(map2,5);
	fmt.Println(map1);
	fmt.Println(map2);
}