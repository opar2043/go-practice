package main
import "fmt"
func main() {
	
	const age = 30;

	if(age >= 40) {
		fmt.Println("he is settled");
	}else if (age < 30) {
		fmt.Println("he is not settled");
	}else {
		fmt.Println("you are kid")
	}

	const a = 1;

	switch a {
	case 1 :
		fmt.Println("age 1")
	case 2 :
		fmt.Println("age 2");
	case 3 : 
	fmt.Println("age 3 ");
	default :
	   fmt.Println("age not here");	
	}

}
