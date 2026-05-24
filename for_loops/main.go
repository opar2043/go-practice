package main

import (
	"fmt" 
)
func main() {
    var numbers[5] int ;
	numbers = [5]int{2,4,6,1,0};

    fmt.Println(numbers[3])

	 for i:=0 ; i < 5; i ++ {
		fmt.Println("Number: ", numbers[i]);
	 }
	 fmt.Println("length :" , len(numbers));
	 fmt.Println("capacity :" , cap(numbers[1:4]));
	 

    for i:=0 ; i < 5 ; i++ {
		if(i == 2) {
			continue
		}
		fmt.Println(i);
	};

	names := []string {"opar" , "rijoan" , "rashid"};
	for i:=0 ; i < len(names) ; i++ {
		fmt.Println(names[i]);
	}
// slice
	mySlice := numbers[1:4];
		for i:=0 ; i < len(mySlice) ; i++ {
		fmt.Println(mySlice[i]);
	}

	// _slice:= numbers[2:];
	slice2 := append(numbers[:],5 , 6);

	fmt.Println(slice2)


	// range
	for idx, value := range mySlice {
		fmt.Println(idx, value);
	}

}