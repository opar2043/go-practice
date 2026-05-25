package main

import (
	"fmt"
	"time"
)
const (
	Red = "red"
	Green = "green"
	Blue = "blue"
)

const (
	deleveried = "deleveried"
	pending = "pending"
	confirmed = "confirmed"
);

func order(status string) {
  fmt.Println("order is ", status)
}

func main(){
	fmt.Println(Red)
	fmt.Println(Green)
	
	// defer will run last 
	defer fmt.Println(Blue)

	status := deleveried
	fmt.Println(status)

	status = pending
	fmt.Println(status)

	status = confirmed
	fmt.Println(status)

	// will run last
	go order(status)

	// will run quickly
	// go order(status)

	time.Sleep(time.Second)
}
