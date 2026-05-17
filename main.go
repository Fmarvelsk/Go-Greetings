package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type person struct {
	name     string
	adddress string
	age      uint8
	height   uint8
	vehicle
}

type vehicle struct {
	brand string
	year  uint16
	color string
}

type recentagle struct {
	length int16
	height int16
}

type circle struct {
	radius int16
}

type calculate interface {
	area() int16
}

func (r recentagle) area() int16 {
	return r.length * r.height
}

func (c circle) area() int16 {
	return int16(2 * 3.14 * float64(c.radius))
}

func shapeArea(s calculate) int16 {
	var result int16 = s.area()

	fmt.Println(result)

	return result
}

func main() {

	var myNumber int // stores interger number only based

	var array [5]int32 = [5]int32{1, 2, 3, 5, 6}

	var newArray = [...]string{"Marvb", "test"} // Arryas

	//var arraySlice = []int32{1, 2, 4} // => Slices (a more powerful)

	b := [...]int{100, 3: 400, 500} /** Addes an element to the array (i.e add 400 to index 3) **/
	fmt.Println("idx:", b)

	var myNumber2 int

	var myString = "Hello world" // type is infer

	myStr := "Next with no var, short-hand"

	test1, test2 := 1, 32

	fmt.Println(test1, test2)

	fmt.Println(myStr, myNumber2, myString, myNumber, array, newArray)
	var addtion, substraction int = operations(5, 8)

	fmt.Printf("Addtion is equal %v while subtraction is %v \n", addtion, substraction)

	var divisionResult, err = devision(50, 10)

	goRoutine()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Division result is %v for all values", divisionResult)
	}
}

func operations(value1 int, value2 int) (int, int) {
	var addResult int = value1 + value2
	var subtractResult int = value1 - value2

	return addResult, subtractResult

}

func devision(value1 int, value2 int) (int, error) {
	var err error

	if value2 == 0 {
		err = errors.New("Value cannot be zero")

		return -1, err
	}

	return value1 / value2, nil

}

func sliceExample() {
	var arraySlice = []int16{20, 50, 30}

	newArraySlice := make([]int16, 5, 8) // make a slice with length 5 and capacity 8

	fmt.Printf("new array slice size is %v and capacity is %v\n", len(newArraySlice), cap(newArraySlice))

	fmt.Printf("array size is %v and capacity is %v\n", len(arraySlice), cap(arraySlice))

	newSlice := []int16{12, 3, 45}

	arraySlice = append(arraySlice, newSlice...)

	fmt.Printf("updated array size is %v and capacity is %v\n", len(arraySlice), cap(arraySlice))

	/** Maps **/
	myMap := make(map[string]uint8)

	setMap := map[string]uint8{"Adam": 40, "Sarah": 60}

	//delete(setMap, "Adam")

	fmt.Println(myMap["age"])

	var value, ok = setMap["Sarah"]

	if ok {
		fmt.Println("Key exist for value", value)
	} else {
		fmt.Println("Key does not exist")
	}

	for i := range setMap {
		fmt.Println("Key is", i, "and value is", setMap[i])
	}

	for key, value := range setMap {
		fmt.Println("Key is", key, "and value is", value)
	}

	item := 0

	for item < 10 {
		fmt.Println("Item is", item)
		item++
	}

	for _, value := range map[string]uint8{"abiola": 2, "yemi": 4, "james": 6} {
		fmt.Println("and value is", value)
	}

	//newString := []rune("r£sume")

	const s = "สวัสดี"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	for i := 0; i < len(s); i++ {
		fmt.Println("Test case key", i, "value is", s[i])
	}

	var buffer [256]byte

	var slice = buffer[100:150]

	fmt.Println(len(slice))

	var sliceString = []string{"Hello", "world", "Go", "programming"}

	var singleString strings.Builder

	for i := range sliceString {
		singleString.WriteString(sliceString[i])
	}

	result := singleString.String()

	fmt.Println(result)
}

func test() {

	myPerson := person{"Thomas", "Lagos", 30, 180, vehicle{"Toyota", 2020, "Red"}}

	fmt.Println(myPerson.brand)

	fmt.Println(myPerson.name)

	var r1 = recentagle{50, 15}
	var c1 = circle{10}

	shapeArea(c1)
	shapeArea(r1)
}

func pointers() {
	var p *int16 = new(int16) // pointers are variable that store memory addess of another variable
	var i int16

	fmt.Printf("Value of p is %v\n", *p)

	fmt.Println("Value of i is", i)

	array := [5]int8{1, 2, 3, 4, 5}

	sum(&array)

	fmt.Printf("Value %p\n", &array)
}

func sum(array *[5]int8) int8 {

	fmt.Printf("Loop %p\n", array)

	var sum int8

	for i := range array {
		sum = +array[i]
	}
	return sum
}

type user struct {
	id       string
	isConfig bool
	name     string
}

var dbModel = map[string]user{"id1": {id: "id1", isConfig: true, name: "james"}, "id2": {id: "id2", isConfig: false, name: "test"}, "id3": {id: "id3", isConfig: false, name: "Joh doe"}, "id4": {id: "id4", isConfig: false, name: "Joh doe"}, "id5": {id: "id5", isConfig: true, name: "Jane"}}
var wg = sync.WaitGroup{}
var recordSlice = []string{}

func dbCall(model user) {

	// simulate a db call
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Printf("user feteched with result %v\n", model)
	recordSlice = append(recordSlice, model.name)

	wg.Done()

}

func goRoutine() {
	t0 := time.Now()
	for i := range dbModel {
		wg.Add(1)
		go dbCall(dbModel[i])
		//recordSlice = append(recordSlice, dbModel[i])
	}
	wg.Wait()
	fmt.Println("Total excuation time is", time.Since(t0))

	fmt.Println("New record slice", recordSlice)

	c := make(chan int8)

	go channels(c)

	for i := range c {

		println("go routine channel", i)
	}

	var chickenChannel = make(chan string)

	var website = [3]string{"https://first.com", "https://second.com", "https://thrid.com"}

	for i := range website {
		wg.Add(1)
		go getUser(website[i], chickenChannel)
	}

	go func() {
		wg.Wait()
		close(chickenChannel)
	}()

	found := false
	for result := range chickenChannel {
		found = true
		println("chicken soup website", result)
	}

	if !found {
		println("no result found")
	}
}

func channels(c chan int8) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- int8(i)
	}
}

func getUser(w string, chickenChannel chan string) {
	defer wg.Done()
	// simulate make api
	time.Sleep(time.Second * 2)
	var price = rand.Float32() * 20

	fmt.Printf("site %v for price %v\n", w, price)
	if price < 5 {
		chickenChannel <- w
	}
}

// generic func
func genric[T int8 | float32](value T, nextValue T) T {

	return value * nextValue
}

// generic can also be used in struct
type mathRule[T circle | recentagle] struct {
	shape T
}

var item = mathRule[circle]{
	shape: circle{
		radius: 30,
	},
}
