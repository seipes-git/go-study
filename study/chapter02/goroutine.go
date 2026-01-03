package main

import (
	"fmt"
	"sync"
	"time"
)

// 基本的goroutine
func sayHello() {
	fmt.Println("hello from goroutine")
}

func useSayHello() {
	go sayHello()
	time.Sleep(time.Second)
}

// 使用WaitGroup等待
func useWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("hello from goroutine", i)
		}(i)
	}

	wg.Wait()
}

// Channel 通信机制
// 无缓冲channel（同步）
func channelDemo() {
	ch := make(chan int)

	go func() {
		ch <- 42 // 发送数据
	}()

	value := <-ch // 接收数据
	fmt.Println(value)
}

// 缓冲channel（异步）
func bufferedChannelDemo() {
	ch := make(chan int, 3) // 缓冲区大小为3

	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 // 这里会阻塞，因为缓冲区满了

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
}

// channel关闭
func closeChannelDemo() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// 关闭后的channel仍然可以读取
	// 不能再发送数据
	for value := range ch {
		fmt.Println(value)
	}
}

// Select 语句
func selectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

// 无缓冲 channel 只要有 “发送方” 和 “接收方” 在不同 goroutine 中同时等待，就能完成数据传递
func timeoutDemo() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case msg := <-ch:
		fmt.Println("收到:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("超时了") // 1秒后输出这个
	}
	// 因为发送需要2秒，但超时是1秒，所以会输出"超时了"
}

func nonBlockingDemo() {
	// ch := make(chan int, 1)
	ch := make(chan int)

	// 非阻塞发送
	select {
	case ch <- 42:
		fmt.Println("发送成功")
	default:
		fmt.Println("channel已满，发送失败")
	}

	// 非阻塞接收
	select {
	case value := <-ch:
		fmt.Println("收到:", value)
	default:
		fmt.Println("没有值可读（非阻塞）") // 立即输出这个
	}

}

func loopSelect() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	// done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- fmt.Sprintf("msg-%d", i)
			time.Sleep(150 * time.Millisecond)
		}
		close(ch2)
	}()

	// 持续监听，直到所有channel都关闭
	for {
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil // 关闭的channel设为nil，select会忽略它
				continue
			}
			fmt.Println("ch1:", val)
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("ch2:", msg)
		default:
			// 如果没有数据，可以做其他事情
			if ch1 == nil && ch2 == nil {
				fmt.Println("所有channel已关闭")
				return
			}
		}
	}
}

func quitChannelDemo() {
	jobs := make(chan int)
	quit := make(chan bool)

	// 工作goroutine
	go func() {
		for {
			select {
			case job := <-jobs:
				fmt.Printf("处理任务: %d\n", job)
			case <-quit:
				fmt.Println("收到退出信号")
				return
			}
		}
	}()

	// 发送任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}

	// 发送退出信号
	quit <- true
	time.Sleep(100 * time.Millisecond)
}

func fairnessDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 同时向两个channel发送数据
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i * 10
		}
		close(ch2)
	}()

	// 持续接收，随机选择就绪的channel
	for i := 0; i < 20; i++ {
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Printf("ch1: %d\n", val)
		case val, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Printf("ch2: %d\n", val)
		}
	}
	// 输出顺序是随机的，体现了公平性
}

// 当channel关闭后：

// 从关闭的channel接收数据，会立即返回零值，ok为false
// 向关闭的channel发送数据，会panic
// 在select中，关闭的channel可以设置为nil，select会忽略它
func closedChannelDemo() {
	ch := make(chan int)
	close(ch)

	select {
	case val, ok := <-ch:
		fmt.Printf("val: %d, ok: %v\n", val, ok) // val: 0, ok: false
	default:
		fmt.Println("default")
	}
}
// func main() {
// 	// useSayHello()
// 	// useWaitGroup()
// 	// channelDemo()
// 	// bufferedChannelDemo()
// 	// closeChannelDemo()
// 	// selectDemo()
// 	// timeoutDemo()
// 	// nonBlockingDemo()
// 	// loopSelect()
// 	// quitChannelDemo()
// 	// fairnessDemo()
// 	// closedChannelDemo()
// }
