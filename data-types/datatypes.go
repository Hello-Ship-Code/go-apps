package main

import "fmt"

func main(){
  name := "Peter"
	var age int = 25 
	b := 20
	c := 12.2
	isLoggedIn := true
	
	fmt.Println("Hello, 👋🏻",name,"your age is",age)
	fmt.Println("Hello, 👋🏻", name, "your age is", age)
	fmt.Printf("Hello, 👋🏻 %v your age is %v.\n",name,age)
	fmt.Println("Addition:",age + b)
	fmt.Println("type casting:",int(c)+b)
	fmt.Println("type casting:",float64(b)+c)


	if(isLoggedIn){
		fmt.Printf("%v is logged in \n",name)
	} else {
		fmt.Println("Please login 🙏🏻")
	}
}