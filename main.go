package main

import (
	"fmt"
)

var b = 300

func main() {

	/* var a int = 5000 // declaration with initialization
	var b = 300      // declaration with initialization
	d := 5000        // declaration with initialization

	// fmt.Println("Enter a Number:  ")
	// fmt.Scanln(&a)
	fmt.Println("The value of a, b and d", a, b, d)
	// fmt.Println("Hello Word")

	// conversion
	var str int = 65
	var e string = strconv.Itoa(str)

	fmt.Printf(" %v, %T ", str, str)
	fmt.Printf(" %v, %T ", e, e)
	*/
	a := 3 + 5i

	fmt.Printf(" %v, %T", a, a)

}

// import (
// 	"github.com/julienschmidt/httprouter"
// )

// func main() {

// 	r := httprouter.New()
// 	uc := controller.NewUserController(getSession)
// 	r.GET("", )
// 	r.POST(" ", )
// 	r.DELETE(" ", )
// 	r.
// }
