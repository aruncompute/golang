package main
import "fmt"

func main(){
	var i int
	var j int

	k:=3
	for i=1; i<=10; i++ {
		j = i << 2;
		fmt.Printf("%d << %d = %d\n",i,k,j);
	}

	//Convert to bitwise operator
	var a1=60;
	var b1 =13;
	bitwiseAdd := a1 & b1
	fmt.Printf("a1 %v=%10b\n",a1,a1)
	fmt.Printf("a1 %v=%10b\n",b1,b1)
	fmt.Printf("Add %v=%b",bitwiseAdd,bitwiseAdd);

 
	
}