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
		if n <= len(queue)-1 {
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
	const USAGE string = "Usage: ./tail [-n #] [file]"
	const NOEXIST string = "No such file or directory"
	intOpt := flag.Int("n", 10, USAGE)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", USAGE)
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
				fmt.Fprintf(os.Stderr, "%s\n", NOEXIST)
				os.Exit(1)
			}
			defer fp.Close()
			show(tail(fp, n))
		}
	} else {
		show(tail(os.Stdin, n))
	}
}
