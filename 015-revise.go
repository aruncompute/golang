package main
import "fmt"

func main(){
	fmt.Printf("\n=========== Exercise 1 =========== \n")
	for x :=1 ; x<=1000; x++{
		fmt.Print(x,",");
	}
	
	fmt.Printf("\n=========== Exercise 2 =========== \n")
	for x:=65; x<=66; x++{
		fmt.Println(x)
		for i :=1; i<=3; i++{
			fmt.Printf("\t%#U\n",x);
		}
		fmt.Print("\n")
	}
	fmt.Printf("\n=========== Exercise 3 =========== \n")
	year :=1984
	for{
		if (year > 2022){
			break;
		}
		fmt.Printf("%v,",year)
		year++;
	}

	fmt.Printf("\n=========== Exercise 4 =========== \n")
	for num:=10 ; num<100; num++{
		fmt.Printf("%v modules 4 = %v\n",num,num%4)
	}

	fmt.Printf("\n=========== Exercise 5 =========== \n")
	//ss5 := "arun"
	switch {
	case false:
		fmt.Println("False")
		break;

	case true:
		fmt.Println("True")
		break;
	default:
		fmt.Println("Default")
	}


	
 
	
}