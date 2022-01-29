package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Admin struct {
	name              string `json: "Field Str"`
	email             string `json: "Field Str"`
	password          string `json: "Field Str"`
	updatedAt         string `json: "Field Str"`
	admin_wallet      int    `json: "Field Int"`
	repurchase_wallet int    `json: "Field Int"`
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://thor:thor@cluster0.ib472.mongodb.net/myfastea_mlm?authSource=admin&replicaSet=atlas-132y7v-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true"))
	if err != nil {
		log.Fatal()
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal()

	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal()
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal()
	}
	fmt.Println(databases)
	col := client.Database("myfastea_mlm").Collection("Admin")
	var admin Admin
	opts := col.FindOne(ctx, admin)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	fmt.Printf("found document %v", opts)
	// clientOptions := options.Client().ApplyURI("mongodb+srv://thor:thor@cluster0.ib472.mongodb.net/myfastea_mlm?authSource=admin&replicaSet=atlas-132y7v-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true")
	// fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	fmt.Println("Mongo.connect() ERROR: ", err)
	// 	os.Exit(1)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// fmt.Println(ctx)

	// col := client.Database("myfastea_mlm").Collection("Admin")
	// fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")
	// collection := Admin
	// cur, err := collection.Find(ctx, Admin)
	// fmt.Printf(cur)
	// var podcast bson.M
	// if err = Admin.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(podcast)
	// oneDoc := MongoField{
	// 	FieldStr:  "This is our first data and its very important",
	// 	FieldInt:  826482746,
	// 	FieldBool: true,
	// }

	// fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	// result, insertErr := col.Find(ctx, oneDoc)
	// if insertErr != nil {
	// 	fmt.Println("InsertONE Error:", insertErr)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
	// 	fmt.Println("InsertOne() api result type: ", result)

	// 	newID := result.InsertedID
	// 	fmt.Println("InsertedOne(), newID", newID)
	// 	fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	// }
}

// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {
// 	// declarationVal()
// 	// hello("Karunendeu")
// 	// goarray()
// 	// gomap()
// 	// createStruct()
// 	// loopInGo()
// 	goPointer()
// }

// func hello(name string) {

// 	var message string
// 	message = fmt.Sprintf("Hi, %v. Welcome!", name)
// 	fmt.Printf(message)

// }

// func goarray() {
// 	var arr [5]int = [5]int{1, 2, 3, 4, 5} //length defin
// 	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
// 	fmt.Println("Array with Length", arr)
// 	fmt.Println("Array without Length", arr1)
// 	fmt.Println("Array Slice ", arr1[:7])

// 	a := 6
// 	// fmt.Printf(" %v, %T", k, k)

// 	pos := sort.SearchInts(arr1[:], a)
// 	fmt.Printf("Found %d at index %d in %v\n", a, pos, arr1)

// 	matrix := [...][]int{{8, 5, 5, 3}, {3, 4, 6}, {9, 8, 7}}
// 	fmt.Println("2'D Array ", matrix, len(matrix))

// 	// sorting
// 	strs := []string{"c", "a", "b"}
// 	sort.Strings(strs)
// 	fmt.Println("Strings:", strs)

// 	ints := []int{7, 2, 4}
// 	sort.Ints(ints)
// 	fmt.Println("Ints:   ", ints)

// 	s := sort.IntsAreSorted(ints)
// 	fmt.Println("Sorted: ", s)

// 	// slicing
// 	var arr2 []int = []int{1, 2, 3, 4, 5}
// 	var arr4 []int = make([]int, 2, 4) // here we declair capicity of array and length
// 	fmt.Println(arr4, len(arr4), cap(arr4))
// 	arr2 = append(arr2, 22)
// 	fmt.Println(arr2, cap(arr2)) // slicing use a address of variable and Array use a value of variable "cap use for capacity"
// }

// func gomap() {
// 	// empSal := make(map[string]int) // decleration
// 	// empSal = map[string]int{       // initialization
// 	// 	"Naha": 20000,
// 	// 	"Raj":  30000,
// 	// 	"Atul": 15000,
// 	// }
// 	empSal := map[string]int{ // initialization
// 		"Naha": 20000,
// 		"Raj":  30000,
// 		"Atul": 15000,
// 	}
// 	fmt.Println(empSal, len(empSal))
// 	fmt.Println(empSal["Raj"]) //get value
// 	empSal["Atul"] = 180000
// 	empSal["Sak"] = 50000 // insert in map

// 	fmt.Println(empSal)

// 	// delete(empSal, "Raj")
// 	fmt.Println(empSal)

// 	//check data ability on map
// 	_, f := empSal["Atul"]
// 	fmt.Println(f) // flag gives a True and false if record found := true , else record not found:= false
// }

// func createStruct() {
// 	type Employee struct {
// 		EmpId     int
// 		EmpName   string
// 		EmpMobile []string
// 	}

// 	emp := Employee{
// 		EmpId:     1001,
// 		EmpName:   "Ravi",
// 		EmpMobile: []string{"1234567890", "79055782623"},
// 	}

// 	fmt.Println(emp)

// 	emp2 := emp
// 	emp.EmpName = "Raga"
// 	fmt.Println(emp)
// 	fmt.Println(emp2)

// }

// func declarationVal() {

// 	var a int = 5000 // declaration with initialization
// 	var b = 300      // declaration with initialization
// 	d := 5000        // declaration with initialization

// 	// fmt.Println("Enter a Number:  ")
// 	// fmt.Scanln(&a)
// 	fmt.Println("The value of a, b and d", a, b, d)
// 	// fmt.Println("Hello Word")

// 	// conversion
// 	var str int = 65
// 	var e string = strconv.Itoa(str)
// 	// k := 3 + 5i

// 	// fmt.Printf(" %v, %T", k, k)
// 	fmt.Printf(" %v, %T ", str, str)
// 	fmt.Printf(" %v, %T ", e, e)
// }

// func loopInGo() {
// 	for i := 0; i < 5; i++ {
// 		if i%2 == 0 {
// 			fmt.Printf("\n %v is Even Number", i)

// 		} else {
// 			fmt.Printf("\n %v is Odd Number", i)

// 		}
// 	}
// }

// func goPointer() {
// 	var x int = 90
// 	var b *int = &x
// 	fmt.Println(x)
// 	fmt.Println(b, *b)

// }

// // import (
// // 	"github.com/julienschmidt/httprouter"
// // )

// // func main() {

// // 	r := httprouter.New()
// // 	uc := controller.NewUserController(getSession)
// // 	r.GET("", )
// // 	r.POST(" ", )
// // 	r.DELETE(" ", )
// // 	r.
// // }
