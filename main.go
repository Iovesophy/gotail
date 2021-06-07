package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var _ tailer = (*fileTail)(nil)
var _ tailer = (*stdinTail)(nil)

type tailer interface {
	appendQueue(*os.File)
	printTail()
}

type stdinTail struct {
	maxQueueSize int
	queue        []string
}

type fileTail struct {
	filename     string
	isNotEndFile bool
	nArg         int
	stdinTail
}

func (s *stdinTail) appendQueue(stream *os.File) {
	defer stream.Close()
	scanner := bufio.NewScanner(stream)
	s.queue = make([]string, s.maxQueueSize, s.maxQueueSize)
	for scanner.Scan() {
		s.queue = append(s.queue, scanner.Text())
		if s.maxQueueSize < len(s.queue) {
			s.queue = s.queue[1:]
		}
	}
}

func (f *fileTail) appendQueue(stream *os.File) {
	f.stdinTail.appendQueue(stream)
}

func (s *stdinTail) printTail() {
	for _, l := range s.queue {
		fmt.Println(l)
	}
}

func (f *fileTail) printTail() {
	if f.nArg > 1 {
		fmt.Printf("==> %s <==\n", f.filename)
	}
	f.stdinTail.printTail()
	if f.isNotEndFile {
		fmt.Println("")
	}
}

func doTail(t tailer, stream *os.File) {
	t.appendQueue(stream)
	t.printTail()
}

func xOpen(filename string) *os.File {
	stream, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return stream
}

func main() {
	const defaultNLines = 10
	nFlags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	nLines := nFlags.Int("n", defaultNLines, "number of lines")
	nFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-n #] [file]\n", os.Args[0])
	}
	nFlags.Parse(os.Args[1:])

	nArg := nFlags.NArg()
	if nArg > 0 {
		for i := 0; i < nArg; i++ {
			t := &fileTail{
				filename:     nFlags.Arg(i),
				isNotEndFile: i+1 < nArg,
				nArg:         nArg,
				stdinTail: stdinTail{
					maxQueueSize: *nLines,
				},
			}
			doTail(t, xOpen(t.filename))
		}
	} else {
		t := &stdinTail{maxQueueSize: *nLines}
		doTail(t, os.Stdin)
	}
}
