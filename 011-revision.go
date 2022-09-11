package main
import "fmt"

func main(){
	//exercise 1
	fmt.Printf("=========== Exercise 1 =========== \n")
 	num := 10;
	fmt.Printf("\tDecimal | Binary | Hexa\n")
 	fmt.Printf("%10d = %8b =%5x\n",num,num,num)

	//excercise 2
	fmt.Printf("\n=========== Exercise 2 =========== \n")
	a1 := 10;
	b1 :=20;

	equalTo := a1==b1
	lessEqualTo := a1 <= b1;
	greaterEqualTo := a1 >= b1;
	notEqualTo := a1 != b1;
	lessThan := a1 < b1;
	greaterThan := a1 > b1;

	fmt.Println(equalTo);
	fmt.Println(lessEqualTo);
	fmt.Println(greaterEqualTo);
	fmt.Println(notEqualTo);
	fmt.Println(lessThan);
	fmt.Println(greaterThan);

	fmt.Printf("\n=========== Exercise 3 =========== \n")
	const(
	a3 = 42
	b3 int = 43
	);
	
	fmt.Println(a3);
	fmt.Println(b3);

	fmt.Printf("\n=========== Exercise 4 =========== \n")
	var a4 int = 8
	fmt.Printf("Decimal | Binary | Hexa\n")
	fmt.Printf("%4d %6b %7x\n",a4,a4,a4)
	a4 = a4 << 1;
	fmt.Printf("Decimal | Binary | Hexa\n")
	fmt.Printf("%5d %8b %7x\n",a4,a4,a4)
	
	fmt.Printf("\n=========== Exercise 5 =========== \n")
	a5 := `this is 'Arun kumar Gupta'
	What is your name?	`
	fmt.Println(a5);
	
	fmt.Printf("\n=========== Exercise 6 =========== \n")
	const a6 int = 10;
	fmt.Println(a6);

	const (
		b6 int iota
	)
	
	fmt.Println(b6);	
	fmt.Println(b6);
	fmt.Println(b6);
 
	
}