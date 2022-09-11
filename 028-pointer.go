package main
import (
	"fmt"	
)

func main(){		
	x := 100;
	y := &x;

	fmt.Printf("X= %v\n",x);
	fmt.Printf("y= %v\n",*y);
	fmt.Println("After change the value of x")
	x=200;
	fmt.Printf("X= %v\n",x);
	fmt.Printf("y= %v\n",*y);

}
