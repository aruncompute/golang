package main
import (
	"fmt"
	"math"
)

type person struct{
	first string
	last string
	age int	
}

func (e person) speak(){
	fmt.Println(e.first, " ", e.last, " Age is ", e.age )
}

type square struct{
	length float64
	
}
type circle struct{
	radius float64

}

func (s square)area() float64{
	return s.length * s.length;
}
func (c circle)area() float64{
	return  math.Pi * c.radius * c.radius;
}

func main(){		
	fmt.Println("================== Excercise 1============");
	defer welcome()

	x1Num := []int{1,2,3,4,5,6,7}
	sumAllNumber :=  sumOfNNumber(x1Num...);
	fmt.Println(sumAllNumber);

	fmt.Println("================== Excercise 2============");
	arun := person{
		first: "Arun",
		last: "Kumar",
		age: 32,
	} 
	arun.speak();

	fmt.Println("================== Excercise 3============");
	c3 := circle{
		radius: 10.5,
	}
	fmt.Println("Area of circle: ",c3.area());
	s3 := square{
		length: 10.5,
	}
	fmt.Println("Area of square: ",s3.area());


}

func welcome(){
	fmt.Println("Welcome to the Golang World!")
}
func sumOfNNumber(number ...int) (int){
	sum := 0;
	for _,val := range(number){
		sum = sum+val;
	}
	
	return sum
}