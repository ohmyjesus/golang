package main

import (
	"context"
	"fmt"
	"time"
)

// context.withdeadline

func main() {
	// 在当前时间上加上50s
	d := time.Now().Add(50 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d) // 时间到期后会向ctx.done里面传值

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
