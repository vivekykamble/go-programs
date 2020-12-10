package main

import "fmt"

type  person struct {
	name string
	age int
	email string
}

func main()  {
	p := person{
		"vivek",
		23,
		"vivekk360@gmail"}
	fmt.Println( p.name, p.age,p.email )

	//Declare and initialize an anonymous struct type with the same three fields. Display the value.
	q:= struct {
		name string
		age int
		email string
	} {
		"kamble",
		33,
		"vivekk360@gmail"}

	fmt.Println( q.name, q.age, q.email )


}