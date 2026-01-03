package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// ## :white_check_mark:指针
// 1. 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
//    - 考察点 ：指针的使用、值传递与引用传递的区别。
// 2. 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
//    - 考察点 ：指针运算、切片操作。

func pointerChangeNum(p *int) int {
	return *p + 10
}

func usePointerChangeNum() {
	a := 10
	var p *int = &a
	fmt.Println(pointerChangeNum(p))
}

// *******************************************************
func sliceChangeNum(s *[]int) {
	for i := 0; i < len(*s); i++ {
		(*s)[i] *= 2
	}
}

func useSliceChangeNum() {
	a := []int{1, 2, 3, 4, 5}
	sliceChangeNum(&a)
	fmt.Println(a)
}

// ## :white_check_mark:Goroutine
// 1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
//    - 考察点 ： go 关键字的使用、协程的并发执行。
// 2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//    - 考察点 ：协程原理、并发任务调度。

func goCounter() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				ch1 <- i
			}
		}
	}()

	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				ch2 <- i
			}
		}
	}()

	for i := 0; i <= 10; i++ {
		select {
		case val := <-ch1:
			fmt.Println("奇数：", val)
		case val := <-ch2:
			fmt.Println("偶数：", val)
		default:
			fmt.Println("没有数据")
		}
	}
}

// **************************************************

type Task func()

type TaskResult struct {
	taskname string
	Duration time.Duration
}

func taskScheduler(tasks map[string]Task) []TaskResult {
	var wg sync.WaitGroup
	ch := make(chan TaskResult, len(tasks))

	for name, task := range tasks {
		wg.Add(1)
		go func(name string, task Task) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			ch <- TaskResult{name, duration}
		}(name, task)
	}

	wg.Wait()
	close(ch)

	var results []TaskResult
	for result := range ch {
		results = append(results, result)
	}

	return results
}

func useTaskScheduler() {
	tasks := map[string]Task{
		"task1": func() {
			time.Sleep(time.Millisecond * 200)
		},
		"task2": func() {
			time.Sleep(time.Millisecond * 100)
		},
		"task3": func() {
			time.Sleep(time.Millisecond * 150)
		},
	}

	fmt.Println("开始执行任务")
	result := taskScheduler(tasks)

	fmt.Println("任务执行详情")
	for _, taskResult := range result {
		fmt.Printf("任务: %s, 耗时: %v\n", taskResult.taskname, taskResult.Duration)
	}
}

// ## :white_check_mark:面向对象
// 1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//    - 考察点 ：接口的定义与实现、面向对象编程风格。
// 2. 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//    - 考察点 ：组合的使用、方法接收者。

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func useInterface() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 20}

	fmt.Printf("circle.area: %.2f\n", circle.Area())
	fmt.Printf("circle.perimeter: %.2f\n", circle.Perimeter())
	fmt.Println("rectangle.area:", rectangle.Area())
	fmt.Println("rectangle.perimeter", rectangle.Perimeter())
}

// *************************************************************
type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func outputInfo() {
	e := Employee{Person{"Tom", 18}, "007"}
	fmt.Println("Name:", e.name, "Age:", e.age, "EmployeeID:", e.EmployeeID)
}

// ## :white_check_mark:Channel
// 1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//   - 考察点 ：通道的基本使用、协程间通信。
//
// 2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//   - 考察点 ：通道的缓冲机制。
func channelDemo() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println(num)
		}
	}()

	wg.Wait()

}

// *******************************************************
func channelDemoWithBuffer() {
	bufCh := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(bufCh)

		for i := 1; i <= 100; i++ {
			bufCh <- i
			fmt.Printf("生产者发送: %d\n", i)
		}
		fmt.Println("生产者已完成所有数据发送")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range bufCh {
			fmt.Printf("消费者接收: %d\n", num)
		}
		fmt.Println("消费者已接收所有数据")
	}()

	wg.Wait()
	fmt.Println("所有操作完成")
}

// ## :white_check_mark:锁机制
// 1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//    - 考察点 ： sync.Mutex 的使用、并发数据安全。
// 2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//    - 考察点 ：原子操作、并发数据安全。

func mutexDemo() {
	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		count int
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

// **************************************************

func atomicDemo() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	finalCount := atomic.LoadInt64(&counter)
	fmt.Printf("最终计数器值：%d\n", finalCount)
}
func main() {
	// usePointerChangeNum()
	// useSliceChangeNum()
	// goCounter()
	// useTaskScheduler()
	// useInterface()
	// outputInfo()
	// channelDemo()
	// channelDemoWithBuffer()
	// mutexDemo()
	atomicDemo()
}
