package main

import (
	"fmt"
	"time"
)

func main() {

	// Создаем канал для передачи строковых данных
	channel := make(chan string)

	// Создаем канал для отправки сигнала остановки
	stop := make(chan bool)

	// Первая горутина, которая отправляет данные в канал
	go func() {
		for {
			select {
			default:
				// Отправляем строку "data" в канал
				channel <- "data"
			case <-stop:
				// При получении сигнала остановки, закрываем канал и выходим из горутины
				close(channel)
				return
			}
		}
	}()

	// Вторая горутина, которая получает данные из канала
	go func() {
		for {
			select {
			default:
				// Получаем данные из канала и выводим их
				fmt.Println(<-channel)
			case <-stop:
				// При получении сигнала остановки, выходим из горутины
				return
			}
		}
	}()

	// Задержка в 1 секунду
	time.Sleep(1 * time.Second)

	// Отправляем сигнал остановки
	stop <- true
}
