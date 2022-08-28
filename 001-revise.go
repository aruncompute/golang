package main
import (
    "fmt"
 //   "time"
)

func main() {	
	//Hello World!
	println("Hello World! => New Line!")
	print("Hello World!\n");
	println("\n===============Variable==============")

	//variable
	var name="Arun kumar Gupta"
	println("My Name is :",name)

	println("\n===========For Loop ==============")
	//for loop
	for i:=1; i<10; i++{
		if i> 5{
			break;
		}
		println("For Loop: ",i);		
	}

	println("\n===========Constant ==============")
	const projectname string = "CONSTANT - Project Name"
	println(projectname);	

	println("\n===========if elseif else ==============")
	var j=10
	if j<10{
		println("j is < 10")
	}else if j>10{
		println("J is > 10")
	} else {
		println("J is 10")
	}
	
	println("\n===========Switch ==============")
	var k=10
	switch k{
	case 1:
		println("Switch - one case")
		break;
		
	case 2:
		println("Switch - Two case")
		break;
	case 3:
		println("Switch - Three case")
		break;
	
	default:
		println("Switch - Default case")
		break;

	}
	
	println("\n===========Array Int|String ==============")
	var list[5] int
	list[0]=100;
	list[1]=101;
	list[2]=102;
	list[3]=103;	
	fmt.Println(list)//normal print will not work

	var arrayPrint[6] string;	
	//arrayPrint := [6]string{"one", "two", "three", "four"}	
	arrayPrint[4] = "Arun kumar Gupta";
	arrayPrint[5] = "Arun ";
	fmt.Println(arrayPrint);

	println("\n===========Map | Associate array ==============")
	associateArray := make(map[string]int)
	associateArray1 := map[string]int{"foo": 1, "bar": 2}

	associateArray["k1"] = 7
    associateArray["k2"] = 13
	associateArray["k3"] = 13
	fmt.Println(associateArray);
	fmt.Println(associateArray1);

	println("\n=========== Range ==============")
	for k,v := range associateArray {
        fmt.Println("key:", k, "\t Value:",v)
        
    }	

     
}