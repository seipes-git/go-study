package main

import (
	"fmt"
	"math"
)

// 接口的定义和实现
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 实现接口（隐式实现）
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 接口作为参数
func PrintArea(s Shape) {
	fmt.Println("面积:", s.Area())
	fmt.Println("周长:", s.Perimeter())
}

func useInterface() {
	r := Rectangle{width: 10, height: 20}
	c := Circle{radius: 5}

	PrintArea(r)
	PrintArea(c)
}

// 空接口可以表示任何类型
var value interface{}

func doSomething(v interface{}) {
	// 方式1：类型断言
	if str, ok := v.(string); ok {
		fmt.Println("字符串:", str)
	}

	// 方式2：type switch
	switch v := v.(type) {
	case string:
		fmt.Println("字符串:", v)
	case int:
		fmt.Println("整数:", v)
	case []int:
		fmt.Println("切片:", v)
	default:
		fmt.Println("未知类型")
	}

}

func useEmptyInterface() {
	// 常见用法：从 map[string]interface{} 中取值并断言具体类型
	payload := map[string]interface{}{
		"id":    1001,
		"name":  "golang",
		"extra": []string{"interface", "assertion"},
	}

	if id, ok := payload["id"].(int); ok {
		fmt.Println("id:", id)
	}

	// 断言失败时要有兜底处理
	if tags, ok := payload["extra"].([]string); ok {
		fmt.Println("tags:", tags)
	} else {
		fmt.Println("extra 字段不是期望的 []string 类型]")
	}

	doSomething(42)
	doSomething("hello")
	doSomething([]int{1, 2, 3})
}

// 接口组合
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	name string
}

func (f *File) Read(p []byte) (int, error) {
	return 0, nil
}

func (f *File) Write(p []byte) (int, error) {
	return 0, nil
}

func useInterfaceCombination() {
	f := &File{name: "file.txt"}

	var r Reader = f
	r.Read(nil)

	var w Writer = f
	w.Write(nil)

	var rw ReadWriter = f
	rw.Read(nil)
	rw.Write(nil)
}
// func main() {
// 	// useInterface()
// 	// useEmptyInterface()
// 	// useInterfaceCombination()
// }
