package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Таймаут
func goroutineWithTimeout(wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("Goroutine with timeout stopped")
			return
		default:
			// Выполнение основной работы
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Закрытие канала
func goroutineWithChannel(wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("Goroutine with channel stopped")
			return
		default:
			// Выполнение основной работы
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Контекст с отменой
func goroutineWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine with context stopped")
			return
		default:
			// Выполнение основной работы
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Флаг завершения
func goroutineWithFlag(stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if *stop {
			fmt.Println("Goroutine with flag stopped")
			return
		}
		// Выполнение основной работы
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// 1. Таймаут
	wg.Add(1)
	go goroutineWithTimeout(&wg)

	// 2. Флаг завершения
	stop := false
	wg.Add(1)
	go goroutineWithFlag(&stop, &wg)

	// Останавливаем через 2 секунды
	time.AfterFunc(2*time.Second, func() {
		stop = true
	})

	// 3. Контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go goroutineWithContext(ctx, &wg)

	// Останавливаем через 2 секунды
	time.AfterFunc(2*time.Second, cancel)

	// 4. Закрытие канала
	done := make(chan struct{})
	wg.Add(1)
	go goroutineWithChannel(&wg, done)

	// Останавливаем через 2 секунды
	time.AfterFunc(2*time.Second, func() {
		close(done)
	})

	wg.Wait()
}
