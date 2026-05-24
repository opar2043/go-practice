package main 
import "fmt"

func swap(num *int)int{
	*num = 20
	return *num
}

func main(){
	x:= 20;
	ptr:= &x;
	fmt.Println("x:" , x);
	// address of the variable
    fmt.Println("adrs of x:" , &x);
    fmt.Println("value of ptr:" , *ptr);
	// access the value
    fmt.Println("adress of ptr:" , &ptr);

	a , b := 6 , 10;

	p1 , p2 := &a , &b;

	// adding with pointer value (*)
	add:= *p1 + *p2;
	fmt.Println("Sum:", add);

	// pointer function swap
    y:= 500
	fmt.Println("Swapped values:");
	// fmt.Println("a:", swap(a));
	fmt.Println("y:", swap(&y));    // 20
	// fmt.Println("b:", swap(b));
}
