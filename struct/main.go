package main
import "fmt"

type User struct{
	name string
	age int
	email string
	isActive bool
}

//dynamically input 

func Input(name string, age int, email string, isActive bool) User {
	return User{
		name:     name,
		age:      age,
		email:    email,
		isActive: isActive,
	}
}

func main(){
	user1 := User{
		name: "John Doe",
		age: 30,
		email: "john.doe@example.com",
		isActive: true,
	}
	fmt.Println(user1)

	user2 := User{
		name: "rijoan",
		age: user1.age,
		email: "rijoan@example.com",
		isActive: false,
	}
	fmt.Println(user2)

	user3 := Input("rashid", 25, "rashid@example.com", true)
	fmt.Println(user3);
}