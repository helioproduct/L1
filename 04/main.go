package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

// Worker принимает канал и идентификатор воркера, считывает элементы из канала и выводит их на экран
func Worker(wg *sync.WaitGroup, source chan interface{}, id int) {
	defer wg.Done() // Уменьшает счетчик WaitGroup на 1 после завершения работы воркера
	for element := range source {
		fmt.Printf("[WORKER %d] received: %s\n", id, element)
	}
}

func main() {
	// Считываем количество воркеров из ввода пользователя
	fmt.Println("input workers count:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSpace(line)

	workersCount, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем канал для передачи данных воркерам
	source := make(chan interface{}, 1)

	// Каналы для обработки сигналов и завершения работы
	sigs := make(chan os.Signal)
	done := make(chan bool)

	// Настраиваем обработку сигнала SIGINT (Ctrl+C)
	signal.Notify(sigs, syscall.SIGINT)

	var wg sync.WaitGroup

	// Запускаем горутину для генерации данных и передачи их в канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(source) // Закрываем канал после получения сигнала завершения
				return
			default:
				source <- "some data" // Отправляем данные в канал
			}
		}
	}()

	// Запускаем воркеры
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go Worker(&wg, source, i)
	}

	<-sigs       // Ожидаем получения сигнала
	done <- true // Отправляем сигнал завершения

	wg.Wait() // Ждем завершения всех горутин
}
