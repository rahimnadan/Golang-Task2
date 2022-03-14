package main 

import "fmt"

func main(){
	//Task2 part 1
	var number uint= 2
	var number_2 float32= 2.5
	var name string = "Abdur Rahim"
	fmt.Println( number, number_2, name)
	//Task 2 part 2
	var i string 
	// Using Scan() function to get values from the user
	fmt.Println("Enter your name")
	fmt.Scan(&i)
	fmt.Println("your name is", i)
	// using Scanln() function to take values from the user
	var j string
	fmt.Println("Enter your father name")
	fmt.Scanln(&j)
	fmt.Println("Your Father name is", j)
	
}
