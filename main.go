package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

func init_queue() ([]string, int) {
	queue := []string{}
	cursor := 0
	return queue, cursor
}

func enqueue(queue []string, value string) []string {
	queue = append(queue, value)
	return queue
}

func dequeue(queue []string) []string {
	queue = queue[1:]
	return queue
}

func show_queue(queue []string, n int) []string {
	if len(queue) == n {
		for i := n; i > 0; i-- {
			if len(queue) != 0 {
				fmt.Println(queue[0])
			}
			queue = dequeue(queue)
		}
	} else {
		for i := len(queue); i > 0; i-- {
			if len(queue) != 0 {
				fmt.Println(queue[0])
			}
			queue = dequeue(queue)
		}
	}
	return queue
}

func tail(stream *os.File, err error, n int) []string {
	queue, cursor := init_queue()
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		if n < 1 {
			n = int(math.Abs(float64(n)))
			if n == 0 {
				n = 10
			}
		}
		queue = enqueue(queue, scanner.Text())
		if n-1 < cursor {
			queue = dequeue(queue)
		}
		cursor++
	}
	return queue
}

func call_tail(stream *os.File, err error, n int) []string {
	queue := tail(stream, err, n)
	queue = show_queue(queue, n)
	return queue
}

func main() {
	const USAGE string = "Usage: gotail [-n #] [file]"
	intOpt := flag.Int("n", 0, USAGE)
	flag.Usage = func() {
		fmt.Println(USAGE)
	}
	flag.Parse()
	if flag.NArg() > 0 {
		fp, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println("Error: No such file or directory")
		}
		defer fp.Close()
		call_tail(fp, err, *intOpt)
	} else {
		call_tail(os.Stdin, nil, *intOpt)
	}
}
