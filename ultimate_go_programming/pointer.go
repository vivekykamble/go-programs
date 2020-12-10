package main

//import "fmt"

func main(){

	//pass by value 
	a:= 10
	
	println("a:\tValue Of[", a, "]\tAddr Of[", &a, "]")
	b:= increament(a)

	println("a:\tValue Of[", a, "]\tAddr Of[", &a, "]")
	println("b:\tValue Of[", b, "]\tAddr Of[", &b, "]")


	// Pass by reference
	var  pointerVariable *int;
	pointerVariable = &a;

	println("Calling Pass by reference")
	increamentByRef(pointerVariable);
	println("a:\tValue Of[", a, "]\tAddr Of[", &a, "]")

	println(" Starting with Struct ")
	p := person{
		"Vivek",
		20,
		"By Construction",
	}
	println(" Before Update Call ")
	println("p:\tValue Of[ person ", p.email, "]\tAddr Of[", &p, "]")
	//var personref = &p
	updateEmail(&p)
	println("after update")
	println("p:\tValue Of[person ", p.email, "]\tAddr Of[", &p, "]")

}

func increament( val int) int{
	val = val + 1
	println("val:\tValue Of[", val, "]\tAddr Of[", &val, "]")
	return val
}

func increamentByRef( val *int) {
	*val++
	println("val:\tValue Of[", *val, "]\tAddr Of[", val, "]")

}

type person struct{
	name string
	age int
	email string
}


func updateEmail(personVar *person){
	//p := *personVar 
	personVar.email = "Set through pass by ref"
	personVar.age =20
}