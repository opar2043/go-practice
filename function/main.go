package main
import "fmt"

func add(n1 int , n2 int)int{
   return n1 + n2;
}

func lan()(string , string , string){
	return "java" , "c++" , "python"
}

// variadic function accepet  miultiple number

func total (numbers ...int)int{
  sum:= 0;
  for _ , value := range numbers{
    fmt.Println("sum:", sum, "value:", value);
    sum= value + sum
    fmt.Println("Recent addition: ", sum)
  }
  return sum
}

func calculate (a int , b int ) (int , int , int , int){
  sum := a + b 
  diff := a - b
  div := a/b;
  mul := a * b;
  return sum , diff , div , mul
}




func  main(){
  res := add(2,5);
  res1, res2, res3, res4 := calculate(20,5);
  // lang := lan()
  fmt.Println(res)
  fmt.Println(res1, res2, res3, res4);


  // annyonomius function
  func(){
  fmt.Println("my name is annyomius function")
  }()


   result := total(2,4,5,7,8,9)
  fmt.Println("Total:", result)
}
