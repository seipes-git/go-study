package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 	Mutex 和 RWMutex
//	使用Mutex保护共享资源

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func useSafeCounter() {
	var c SafeCounter
	for i := 0; i < 100; i++ {
		go c.Inc()
	}
	fmt.Println(c.Value())
}

// WaitGroup同步
func WaitGroup() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	sum := 0

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			time.Sleep(100 * time.Millisecond)

			mu.Lock()
			sum += id
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(sum)
}

// Context 上下文控制

func cancellableDemo() {
	fmt.Println("=== 可取消的Context ===")

	ctx, cancel := context.WithCancel(context.Background())

	// 启动工作goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("工作goroutine收到取消信号:", ctx.Err())
				return
			default:
				fmt.Println("工作goroutine正在运行")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("发送取消信号")
	cancel()

	// 等待goroutine退出 case <-ctx.Done():
	time.Sleep(500 * time.Millisecond)
}

func timeoutContextDemo() {
	fmt.Println("=== 超时Context ===")

	// 设置1秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan string)

	// 模拟一个耗时2秒的操作
	go func() {
		//time.Sleep(200 * time.Millisecond)
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case <-ctx.Done():
		fmt.Println("超时了:", ctx.Err())
	case msg := <-ch:
		fmt.Println("收到:", msg)
	}
}

// 截止时间示例
func deadlineContextDemo() {
	fmt.Println("=== 截止时间Context ===")

	// 设置3秒后的截止时间
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// 检查剩余时间
	if d, ok := ctx.Deadline(); ok {
		fmt.Printf("截止时间: %v, 剩余: %v\n", d, time.Until(d))
	}

	// 等待超过截止时间
	time.Sleep(4 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("截止时间已过:", ctx.Err())
	default:
		fmt.Println("截止时间未过")
	}
}

// 传递值示例
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
)

func valueContextDemo() {
	fmt.Println("=== Context传递值 ===")

	// 创建携带值的context
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-123")
	ctx = context.WithValue(ctx, userIDKey, "user-456")

	// 在函数中读取值
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	if reqID := ctx.Value(requestIDKey); reqID != nil {
		fmt.Printf("Request ID: %v\n", reqID)
	}

	if userID := ctx.Value(userIDKey); userID != nil {
		fmt.Printf("User ID: %v\n", userID)
	}
}

// Context的高级使用场景
// 场景1：级联取消（父取消，子也取消）
func cascadeCancelDemo() {
	fmt.Println("=== 级联取消示例 ===")

	// 创建父context
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	// 创建子context
	childCtx1, cancel1 := context.WithCancel(parentCtx)
	defer cancel1()

	childCtx2, cancel2 := context.WithCancel(parentCtx)
	defer cancel2()

	// 启动子goroutine
	go worker(childCtx1, "Worker 1")
	go worker(childCtx2, "Worker 2")

	time.Sleep(1 * time.Second)

	// 取消父context，所有子context也会被取消
	fmt.Println("取消父context")
	parentCancel()

	time.Sleep(500 * time.Millisecond)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 收到取消信号\n", name)
			return
		default:
			fmt.Printf("%s: 工作中...\n", name)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// 场景2：HTTP超时处理
func httpRequestDemo() {
	// 创建5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, "GET", "https://github.com/", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("请求超时")
		} else {
			fmt.Println("请求失败:", err)
		}
		return
	}
	defer resp.Body.Close()

	fmt.Println("请求成功:", resp.StatusCode)
}

// 场景3：数据库查询超时处理
func databaseQueryDemo(db *sql.DB) {
	// 创建3秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 执行查询
	rows, err := db.QueryContext(ctx, "SELECT * FROM users WHERE active = ?", true)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("数据库查询超时")
		} else {
			fmt.Println("查询失败:", err)
		}
		return
	}
	defer rows.Close()

	// 处理结果
	for rows.Next() {
		// ...
		fmt.Println("处理结果...")
	}
}

// 场景4：多worker协同工作
func multiWorkerDemo() {
	fmt.Println("=== 多worker协同工作 ===")

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// 启动多个worker
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d: 退出\n", id)
					return
				default:
					fmt.Printf("Worker %d: 处理任务\n", id)
					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i)
	}

	// 工作2秒后取消所有worker
	time.Sleep(2 * time.Second)
	fmt.Println("发送取消信号给所有worker")
	cancel()

	// 等待所有worker退出
	wg.Wait()
	fmt.Println("所有worker已退出")
}

// 场景5：Context在Pipeline中的应用
func pipelineDemo() {
	fmt.Println("=== Pipeline示例 ===")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Stage 1: 生成数据
	dataCh := generateData(ctx)

	// Stage 2: 处理数据
	processedCh := processData(ctx, dataCh)

	// Stage 3: 输出结果
	for result := range processedCh {
		fmt.Println("最终结果:", result)
	}

	fmt.Println("Pipeline完成")
}

func generateData(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("生成器: 收到取消信号")
				return
			case ch <- i:
				fmt.Println("生成器: 生成", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	return ch
}

func processData(ctx context.Context, input <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for data := range input {
			select {
			case <-ctx.Done():
				fmt.Println("处理器: 收到取消信号")
				return
			case ch <- data * 2:
				fmt.Println("处理器: 处理", data, "->", data*2)
			}
		}
	}()
	return ch
}
func main() {
	// useSafeCounter()
	// WaitGroup()
	// cancellableDemo()
	// timeoutContextDemo()
	// deadlineContextDemo()
	// valueContextDemo()
	// cascadeCancelDemo()
	// httpRequestDemo()
	// databaseQueryDemo()
	// multiWorkerDemo()
	// pipelineDemo()
}
