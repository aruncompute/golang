package main
import "fmt"

func main(){

	fmt.Printf("\n=========== Exercise 1 =========== \n")
	x1 := [5]int{1,2,3,4,5}

	for i, v :=range x1{
		fmt.Println(i,v);
	}
	fmt.Printf("\n=========== Exercise 2 =========== \n")
	x2 := []int{1,2,3,4,5}		
	x3 := append(x2,11);	
	fmt.Println(x2);
	fmt.Println(x3);

	fmt.Printf("\n=========== Exercise 4 =========== \n")
	var x4 = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51};
	fmt.Println(x4[:5]);
	fmt.Println(x4[5:10]);
	fmt.Println(x4[2:7]);
	fmt.Println(x4[1:6]);

	fmt.Printf("\n=========== Exercise 5 =========== \n")
	x5 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x5 = append(x5,52);
	fmt.Println(x5);

	x5 = append(x5,53,54,55);
	fmt.Println(x5);

	y5 := []int{56, 57, 58, 59, 60}
	x5 = append(x5,y5...);
	fmt.Println(x5);

	fmt.Printf("\n=========== Exercise 6 =========== \n")
	x6 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y6	:= append(x6[0:3],x6[6:10]...);	
	fmt.Println(y6);
	
	fmt.Printf("\n=========== Exercise 7 =========== \n")
	x7 := [][]string{{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(x7);
	for i, nameArray := range(x7){		
		for j, value :=range(nameArray){
			fmt.Println(i,j,value);
		}
	}
	
	fmt.Printf("\n=========== Exercise 8 =========== \n")
	x8 := map[string][]string{
		"name1": []string{"James", "Bond", "Shaken, not stirred"},
		"name2": []string{"Miss", "Moneypenny", "Helloooooo, James."}}

		delete(x8,"name2");
		for i, nameArray := range(x8){		
			for j, value :=range(nameArray){
				fmt.Println(i,j,value);
			}
		}

}