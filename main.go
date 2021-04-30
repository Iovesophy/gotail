package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

func tail(stream *os.File, err error, n int) []string {
	queue := []string{}
	cursor := 0
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		queue = append(queue, scanner.Text())
		if n <= cursor {
			queue = queue[1:]
		}
		cursor++
	}
	ex_queue := queue
	for i := len(queue); i > 0; i-- {
		if len(queue) != 0 {
			fmt.Println(queue[0])
		}
		queue = queue[1:]
	}
	return ex_queue
}

func main() {
	var fp *os.File
	var err error
	const USAGE string = "Usage: gotail [-n #] [file]"
	intOpt := flag.Int("n", 10, USAGE)
	flag.Usage = func() {
		fmt.Println(USAGE)
	}
	flag.Parse()
	if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			if i > 0 {
				fmt.Print("\n")
			}
			if flag.NArg() != 1 {
				fmt.Println("==> " + flag.Arg(i) + " <==")
			}
			fp, err = os.Open(flag.Arg(i))
			if err != nil {
				fmt.Println("Error: No such file or directory")
				os.Exit(1)
			}
			defer fp.Close()
			tail(fp, err, int(math.Abs(float64(*intOpt))))
		}
	} else {
		tail(os.Stdin, nil, int(math.Abs(float64(*intOpt))))
	}
}
