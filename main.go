package main

import (
	"bufio"
	"fmt"
	"os"
)

func enqueue(queue []string, value string) []string {
	queue = append(queue, value)
	return queue
}

func dequeue(queue []string) ([]string, string) {
	value := queue[0]
	queue = queue[1:]
	return queue, value
}

func tail(n int) {
	queue := []string{}
	cursor := 0
	n_lines := 100
	fmt.Println(n_lines)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if n_lines-n-1 < cursor {
			queue = enqueue(queue, scanner.Text())
		}
		cursor++
	}

	fmt.Println("check queue")
	fmt.Println(queue)

	var value string
	for i := n; i > 0; i-- {
		queue, value = dequeue(queue)
		fmt.Println(value)
	}

	fmt.Println("check queue")
	fmt.Println(queue)
}

func main() {
	tail(5)
}
