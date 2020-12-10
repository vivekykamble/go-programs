package main
import "fmt"
func main()  {
	var a =2
	var  b int32 
	var c string
	var d bool
	fmt.Println( a,b,c,d )

	e := 4
	 f := false
	 g :="Stirngf"
	 h :=3.4

	fmt.Println( e,f,g,h )

	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)
}