package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func tail(stream *os.File, n int) []string {
	queue := []string{}
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		queue = append(queue, scanner.Text())
		if n < len(queue) {
			queue = queue[1:]
		}
	}
	return queue
}

func show(queues []string) {
	for _, queue := range queues {
		fmt.Println(queue)
	}
}

func main() {
	const USAGE string = "usage: ./tail [-n #] [file ...]"
	intOpt := flag.Int("n", 10, USAGE)
	flag.Usage = func() {
		err := fmt.Errorf("tail: %s", USAGE)
		println(err.Error())
	}
	flag.Parse()
	n := *intOpt
	if flag.NArg() != 0 {
		for i := 0; i < flag.NArg(); i++ {
			if flag.NArg() != 1 {
				if i != 0 {
					fmt.Print("\n")
				}
				fmt.Printf("==> %s <==\n", flag.Arg(i))
			}
			fp, err := os.Open(flag.Arg(i))
			if err != nil {
				err := fmt.Errorf("tail: %s", err)
				println(err.Error())
				os.Exit(1)
			}
			defer fp.Close()
			show(tail(fp, n))
		}
	} else {
		show(tail(os.Stdin, n))
	}
}
