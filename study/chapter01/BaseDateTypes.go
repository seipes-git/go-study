package chapter01

import (
	"fmt"
	"strconv"
)

func boolType() {
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println(isTrue, isFalse)
}

func intType() {
	var a int = -10
	var b int8 = 10
	var c int16 = -10
	var d int32 = 10
	var e int64 = 10

	fmt.Println(a, b, c, d, e)
}

func uintType() {
	var a uint = 10
	var b uint8 = 10
	var c uint16 = 10
	var d uint32 = 10
	var e uint64 = 10

	fmt.Println(a, b, c, d, e)
}

func floatType() {
	var a float32 = 10.10
	var b float64 = 10.10

	fmt.Println(a, b)
}

func typeChange() {
	var a int = 10
	var b float32 = float32(a)

	var c float64 = 10.10
	var d int = int(c)

	var str string = "321" //"hello world"
	var num, err = strconv.Atoi(str)

	fmt.Println(b, d, num, err)
}
func stringType() {
	var a string = "hello world"
	var name = "Alan"

	str := "Niceday"
	firstByte := str[0]

	raw := `没问题
	hello world
	！！！
	`

	str1 := "come"
	str2 := "on baby"
	combine := str1 + " " + str2

	fmt.Println(a)
	fmt.Println(name)
	fmt.Println(firstByte)
	fmt.Println(raw)
	fmt.Println(combine)
}

func pointerType() {
	var a int = 10
	var b *int
	b = &a

	fmt.Println(a, *b, b)
}
func arrayType() {
	var a [5]int
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5}

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	fmt.Println(a, arr, arr2)
}
func sliceType() {
	var slice []int              // empty slice
	slice = []int{1, 2, 3, 4, 5} // init slice
	slice6 := make([]int, 3, 5)  // make([]Type, length, capacity)

	slice1 := append(slice, 6)
	slice2 := slice[1:4]
	slice3 := slice[:3]
	slice4 := slice[1:]
	slice5 := slice[:]

	length := len(slice)
	capacity := cap(slice1)

	fmt.Println(slice, slice1, slice2, slice3, slice4, slice5, slice6, length, capacity)
}

func mapType() {
	var m map[string]int
	m = make(map[string]int)
	m2 := map[string]int{
		"hello": 1,
		"world": 2,
	}

	m["Alan"] = 24
	m["John"] = 25
	value := m["Alan"]
	delete(m, "Mike") // if delete fail , nothing will happen
	value2, ok := m["Mike"]

	delete(m2, "hello")

	fmt.Println(m, m2, value, value2, ok)

	for key, value := range m {
		fmt.Println(key, value)
	}

	for _, value := range m {
		fmt.Println(value)
	}

	for key := range m {
		fmt.Println(key)
	}
}

// func main() {

// 	// boolType()
// 	// intType()
// 	// uintType()
// 	// floatType()
// 	// typeChange()
// 	// stringType()
// 	// pointerType()
// 	// arrayType()
// 	// sliceType()
// 	// mapType()
// }
