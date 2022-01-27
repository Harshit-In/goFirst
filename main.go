package main

import (
	"fmt"
	"sort"
)

func main() {

	hello("Karunendeu")
	goarray()
	gomap()
}

func hello(name string) {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Printf(message)

}

func goarray() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5} //length defin
	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Array with Length", arr)
	fmt.Println("Array without Length", arr1)
	fmt.Println("Array Slice ", arr1[3:])

	a := 6
	pos := sort.SearchInts(arr1[:], a)
	fmt.Printf("Found %d at index %d in %v\n", a, pos, arr1)

	matrix := [...][]int{{8, 5, 3}, {3, 4, 6}, {9, 8, 7}}
	fmt.Println("2'D Array ", matrix, len(matrix))
}

func gomap() {
	// empSal := make(map[string]int) // decleration
	// empSal = map[string]int{       // initialization
	// 	"Naha": 20000,
	// 	"Raj":  30000,
	// 	"Atul": 15000,
	// }
	empSal := map[string]int{ // initialization
		"Naha": 20000,
		"Raj":  30000,
		"Atul": 15000,
	}
	fmt.Println(empSal, len(empSal))
	fmt.Println(empSal["Raj"]) //get value
	empSal["Atul"] = 180000
	empSal["Sak"] = 50000 // insert in map

	fmt.Println(empSal)
}

/* func main() {

	var a int = 5000 // declaration with initialization
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

	a := 3 + 5i

	fmt.Printf(" %v, %T", a, a)

} */

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
