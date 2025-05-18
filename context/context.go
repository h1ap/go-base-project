package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// TestWithCancel()
	// TestWithTimeout()
	// TestWithValue()
	TestWithDeadline()
	time.Sleep(10 * time.Second)
}

func TestWithDeadline() {
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(1*time.Second))
	go buyFlowers(ctx)
	go buyOil(ctx)
	go buyEggs(ctx)
}

func TestWithValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name1", "Daddy")
	go goToDaddy(ctx)
}

func goToDaddy(ctx context.Context) {
	name := ctx.Value("name1")
	fmt.Println("go to daddy", name)
	ctx = context.WithValue(ctx, "name2", "Mom")
	go goToMom(ctx)
}

func goToMom(ctx context.Context) {
	name := ctx.Value("name2")
	fmt.Println("go to mom", name)
	ctx = context.WithValue(ctx, "name3", "Son")
	go goToSon(ctx)
}

func goToSon(ctx context.Context) {
	name := ctx.Value("name1")
	fmt.Println("go to son", name)
	name = ctx.Value("name2")
	fmt.Println("go to son", name)
	name = ctx.Value("name3")
	fmt.Println("go to son", name)
}

func TestWithTimeout() {
	ctx := context.Background()
	// 自动 cancel
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)

	go buyFlowers(ctx)
	go buyOil(ctx)
	go buyEggs(ctx)
}

func TestWithCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// fmt.Println(ctx.Err())

	go buyFlowers(ctx)
	go buyOil(ctx)
	go buyEggs(ctx)

	// 该步骤会超时
	// time.Sleep(3 * time.Second)
	// cancel()

	// 该步骤会提前取消
	time.Sleep(3 * time.Second)
	cancel()
}

func buyFlowers(ctx context.Context) {
	fmt.Println("buy flowers start")
	select {
	case <-ctx.Done():
		fmt.Println("buy flowers done")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("buy flowers timeout")
		return
	}
}

func buyOil(ctx context.Context) {
	// context 具有继承传播性，子 context 可以访问父 context，反之不行
	ctx, _ = context.WithCancel(ctx)
	fmt.Println("buy oil start")
	select {
	case <-ctx.Done():
		fmt.Println("buy oil done")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("buy oil timeout")
		return
	default:
		fmt.Println("buy oil default")
	}

	go buyEggs(ctx)
}

func buyEggs(ctx context.Context) {
	fmt.Println("buy eggs start")
	select {
	case <-ctx.Done():
		fmt.Println("buy eggs done")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("buy eggs timeout")
		return
	}
}
