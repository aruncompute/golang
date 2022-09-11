package main
import "fmt"

func main(){	
	fun1();

	fun2("Arun Kumar");
	fun2();
	
	sum := fun3Sum(10,20)
	fmt.Println("Sum of two Number: ",sum);

	sum1, mul1 := fun4SumMulti(10,20);
	fmt.Println("Sum of two Number: ",sum1);
	fmt.Println("Multiply of two Number: ",mul1);
}

func fun1(){
	fmt.Println("Hello World");
}
func fun2(name ...string){
	fmt.Println("Welcome to Go lang: ",name);
}
func fun3Sum(a int, b int) int{
	return a+b;
}

func fun4SumMulti(a int, b int) (int, int){
	sum := a+b;
	mul := a*b;
	return sum , mul
}