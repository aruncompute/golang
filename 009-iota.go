package main
import "fmt"

func main(){
	const(
		a = iota
		b = iota
		c = iota
	)	
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T, %T, %T",a,b,c,)	

	fmt.Println(a,b,c)
 
	
}