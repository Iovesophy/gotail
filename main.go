package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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
	for queue := range queues {
		fmt.Println(queue)
	}
}

func main() {
	const USAGE string = "Usage: gotail [-n #] [file]"
	intOpt := flag.Int("n", 10, USAGE)
	flag.Usage = func() {
		fmt.Println(USAGE)
	}
	flag.Parse()
	n := int(math.Abs(float64(*intOpt)))
	if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			if i > 0 {
				fmt.Print("\n")
			}
			if flag.NArg() != 1 {
				fmt.Println("==> " + flag.Arg(i) + " <==")
			}
			fp, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Println("Error: No such file or directory")
				os.Exit(1)
			}
			defer fp.Close()
			show(tail(fp, n))
		}
	} else {
		show(tail(os.Stdin, n))
	}
}
