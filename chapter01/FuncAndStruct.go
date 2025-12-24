package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}

func calculate(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func useCalculate() {
	fmt.Println(add(5, 3))
	fmt.Println(multiply(5, 3))
	fmt.Println(divide(5, 3))
	fmt.Println(calculate(5, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

// 闭包（closure）的核心特征：
// 函数嵌套：闭包是定义在函数内部的函数
// 捕获外部变量：内部函数可以访问外部函数的局部变量
// 变量持久化：即使外部函数执行完毕，内部函数仍然可以访问这些变量
// 每次调用独立：每次调用外部函数都会创建一个新的闭包实例

func makeCounter() func() int {
	count := 0

	// 内部函数 (闭包)
	return func() int {
		count++
		return count
	}
}

func useCounter() {
	counter := makeCounter() // 创建闭包实例
	fmt.Println(counter())   // 计数跟随实例
	fmt.Println(counter())
	fmt.Println(counter())

	counter2 := makeCounter() // 创建第二个闭包实例
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter())
}

// 闭包的高级用法：柯里化（Currying）
// 柯里化是一种将接受多个参数的函数转换为一系列接受单个参数的函数的技术。

// addCurry 接收一个整数参数 a，返回值类型是 func(int) int（即 “接收整数、返回整数” 的函数）。
// 简单说：它不是直接算结果，而是返回一个 “能算结果的工具”。
func addCurry(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
func useCurry() {
	add5 := addCurry(5)  // 5 + b
	fmt.Println(add5(3)) // 5 + 3

	result := addCurry(5)(3)
	fmt.Println(result)
}

// 闭包的实际应用场景
// 1、函数工厂（创建配置化的函数）

func createLogger(prefix string) func(string) {
	return func(message string) {
		//fmt.Println("[%s] %s\n", prefix, message)
		// 使用 fmt.Println 会将 "[%s] %s\n" 看作一个字符串
		fmt.Printf("[%s] %s\n", prefix, message)
	}
}

func useLogger() {
	infoLogger := createLogger("INFO")
	infoLogger("This is an info message")

	errorLogger := createLogger("ERROR")
	errorLogger("This is an error message")
}

// 2、封装私有状态
func createAccount(initialBalance int) (deposit, withdraw, getBalance func(int) int) {
	balance := initialBalance // 私有变量，外部无法直接访问

	// 返回的函数名（deposit/withdraw/getBalance）是 “容器”，闭包函数是 “装进容器里的内容”
	// 容器的名字和内容（闭包）本身无关，只是我们用和功能匹配的名字（deposit 对应存款逻辑）让代码更易读。
	deposit = func(amount int) int {
		balance += amount
		return balance
	}

	withdraw = func(amount int) int {
		if amount <= balance {
			balance -= amount
		}
		return balance
	}

	getBalance = func(_ int) int {
		return balance
	}

	return
}

// 函数返回时，会严格按照 (deposit, withdraw, getBalance) 的顺序输出三个函数值：
// 第一个返回值：永远是 “存款闭包”；
// 第二个返回值：永远是 “取款闭包”；
// 第三个返回值：永远是 “查余额闭包”。
// 调用时，你声明的变量会按顺序绑定这三个返回值：
func useAccount() {
	// 使用
	deposit, withdraw, getBalance := createAccount(100)
	fmt.Println(deposit(50))   // 余额: 150
	fmt.Println(withdraw(30))  // 余额: 120
	fmt.Println(getBalance(0)) // 输出: 120
	// balance 变量无法从外部直接访问，实现了数据封装
}

// 3、延迟执行和惰性求值
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func useFibonacci() {
	fib := fibonacci()
	fmt.Println(fib()) // 输出: 0
	fmt.Println(fib()) // 输出: 1
	fmt.Println(fib()) // 输出: 1
	fmt.Println(fib()) // 输出: 2
	fmt.Println(fib()) // 输出: 3
}

// 4、回调函数和事件处理
func createHandler(name string) func() {
	count := 0
	return func() {
		count++
		fmt.Printf("[%s] Called %d times\n", name, count)
	}
}

func useHandler() {
	handler := createHandler("Click")
	handler()
	handler()
	handler()
}

// ---------------------------STRUCT----------------------------
type Person struct {
	Name string
	Age  int
}

var p1 = Person{
	Name: "Alice",
	Age:  30,
}

var p2 = Person{"Bob", 32}

var p3 = Person{Name: "Charlie"}

// 定义方法（值接收者）
func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// 定义方法（指针接收者，可以修改结构体）
func (p *Person) SetAge(age int) {
	p.Age = age
}

// 全局变量需要使用var进行声明， := 为短变量声明，只可用于局部变量
func useStruct() {
	fmt.Println(p1.GetInfo()) // 输出: Name: Alice, Age: 30
	p2.SetAge(35)
	fmt.Println(p2.GetInfo()) // 输出: Name: Bob, Age: 35
}

// 嵌入结构体 (类似继承)
type Empolyee struct {
	Person
	ID string
}

func (e Empolyee) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, ID: %s", e.Name, e.Age, e.ID)
}

func useEmployee() {
	e := Empolyee{Person{"Alice", 30}, "E001"}
	fmt.Println(e.GetInfo()) // 输出: Name: Alice, Age: 30, ID: E001
}

func main() {

	// useCalculate()
	// useCounter()
	// useCurry()
	// useLogger()
	// useAccount()
	// useFibonacci()
	// useHandler()
	useStruct()
	useEmployee()
}
