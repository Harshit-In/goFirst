package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// declarationVal()
	// hello("Karunendeu")
	// goarray()
	// gomap()
	// createStruct()
	loopInGo()
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
	fmt.Println("Array Slice ", arr1[:7])

	a := 6
	// fmt.Printf(" %v, %T", k, k)

	pos := sort.SearchInts(arr1[:], a)
	fmt.Printf("Found %d at index %d in %v\n", a, pos, arr1)

	matrix := [...][]int{{8, 5, 5, 3}, {3, 4, 6}, {9, 8, 7}}
	fmt.Println("2'D Array ", matrix, len(matrix))

	// sorting
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// slicing
	var arr2 []int = []int{1, 2, 3, 4, 5}
	var arr4 []int = make([]int, 2, 4) // here we declair capicity of array and length
	fmt.Println(arr4, len(arr4), cap(arr4))
	arr2 = append(arr2, 22)
	fmt.Println(arr2, cap(arr2)) // slicing use a address of variable and Array use a value of variable "cap use for capacity"
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

	// delete(empSal, "Raj")
	fmt.Println(empSal)

	//check data ability on map
	_, f := empSal["Atul"]
	fmt.Println(f) // flag gives a True and false if record found := true , else record not found:= false
}

func createStruct() {
	type Employee struct {
		EmpId     int
		EmpName   string
		EmpMobile []string
	}

	emp := Employee{
		EmpId:     1001,
		EmpName:   "Ravi",
		EmpMobile: []string{"1234567890", "79055782623"},
	}

	fmt.Println(emp)

	emp2 := emp
	emp.EmpName = "Raga"
	fmt.Println(emp)
	fmt.Println(emp2)

}

func declarationVal() {

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
	// k := 3 + 5i

	// fmt.Printf(" %v, %T", k, k)
	fmt.Printf(" %v, %T ", str, str)
	fmt.Printf(" %v, %T ", e, e)
}

func loopInGo() {

	for i := 0; i < 5; i++ {
		fmt.Printf("\n*****")
	}
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
