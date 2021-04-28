package main

import (
	"bufio"
	"fmt"
	"os"
)

func init_queue() ([]string, int) {
	queue := []string{}
	cursor := 0
	return queue, cursor
}

func check_queue(queue []string) {
	fmt.Println("check queue")
	fmt.Println(queue)
}

func enqueue(queue []string, value string) []string {
	queue = append(queue, value)
	return queue
}

func dequeue(queue []string) []string {
	queue = queue[1:]
	return queue
}

func show_queue(queue []string, n int) {
	for i := n; i > 0; i-- {
		value := queue[0]
		fmt.Println(value)
		queue = dequeue(queue)
	}
}

func tail(stream *os.File, n int) {
	queue, cursor := init_queue()
	check_queue(queue)

	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		queue = enqueue(queue, scanner.Text())
		if n-1 < cursor {
			queue = dequeue(queue)
		}
		cursor++
	}
	check_queue(queue)

	show_queue(queue, n)

	check_queue(queue)
}

func main() {
	stream := os.Stdin
	tail(stream, 5)
}
