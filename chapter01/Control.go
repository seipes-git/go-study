package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func ifelseType() {

	age := 16
	score := 80

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}

	if num := 10; num > 8 {
		fmt.Println("num is bigger than 8")
	}

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}

func switchType() {

	day := 5

	// 默认不穿透，不需要使用break
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	//	多值匹配
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Working day")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	var score int = 90

	switch {
	case score >= 90:
		fmt.Println("more than 90")
		fallthrough
	case score >= 80:
		fmt.Println("more than 80")
		fallthrough
	case score >= 70:
		fmt.Println("more than 70")
		fallthrough
	default:
		fmt.Println("default")
	}
}

func forType() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	arr := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	str := "hello world 美好的一天"
	for index, char := range str {
		fmt.Println(index, char)
	}

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

//-------------------defer type--------------------

// defer 在函数返回之前执行
// 执行完 return 语句之后，在函数真正返回之前
func example() int {
	defer fmt.Println("world")
	fmt.Println("hello")
	return 42
}

// defer 在 panic 之后也会执行
func panicExample() {
	defer fmt.Println("mission complete")
	fmt.Println("mission start")
	panic("something bad happened")
	fmt.Println("this will not be printed")
}

// defer 执行顺序：LIFO（后进先出，Last In First Out）
// 类似于栈
func multiplyDefer() {
	fmt.Println("start")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("end")
}

// defer 捕获变量值的时机（立即计算）
// defer 语句中的参数值会在 defer 语句执行时立即计算并捕获，而不是在 defer 真正执行时计算
func deferValue() {
	i := 0
	defer fmt.Println("defer 1:", i)
	i++
	defer fmt.Println("defer 2:", i)
	i++
	fmt.Println("i =", i)
	return
}

// defer 的常用应用场景
// 1、释放资源

func resourceRelease() error {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close() // 确保无论函数运行成功与否，文件都会被关闭

	return nil
}

// 2、解锁互斥锁
func safeOperation() {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	// 互斥锁被释放
}

// 3、回复panic（配合recover使用）
func recoverPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	panic("Something bad happened")
}

// 4、性能测量
func performanceMeasurement() {
	defer func() {
		fmt.Println("Finished execution")
	}()

	startTime := time.Now()

	defer func() {
		elapsed := time.Since(startTime)
		fmt.Println("Execution time:", elapsed)
	}()

	// 执行耗时较长的代码
	deferValue()
}
func deferType() {
	// result := example()
	// fmt.Println("返回值:", result)

	// panicExample()

	// multiplyDefer()

	// deferValue()

	// resourceRelease()

	// safeOperation()

	// recoverPanic()

	// performanceMeasurement()
}

func main() {
	// ifelseType()
	// switchType()
	// forType()
	deferType()
}
