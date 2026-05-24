package main
import "fmt"

func main(){
	users := map[string]int{"opar" : 25 , "rijoan" : 30}
	fmt.Println("value:", users["opar"])

	country:=map[string ]string{"BD " : "Bangladesh" , "IND" : "India" , "USA" : "United States" , "UK" : "United Kingdom"};

	for key,value := range country{
		fmt.Println("Key:", key, "Value:", value)
	}
	
	for _,value := range country{
		fmt.Println( "Value:", value)
	}

	for k,v := range  "opar" {
		fmt.Println("idx:", k, "ascii value:", v , "add:" , v + v)
	}


}
