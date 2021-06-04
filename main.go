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
	queueData    []string
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
	for scanner.Scan() {
		s.queueData = append(s.queueData, scanner.Text())
		if s.maxQueueSize < len(s.queueData) {
			s.queueData = s.queueData[1:]
		}
	}
}

func (f *fileTail) appendQueue(stream *os.File) {
	f.stdinTail.appendQueue(stream)
}

func (s *stdinTail) printTail() {
	for _, l := range s.queueData {
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
		for i, j := 0, 1; i < nArg; i, j = i+1, j+1 {
			t := new(fileTail)
			t.filename = nFlags.Arg(i)
			t.isNotEndFile = j < nArg
			t.nArg = nArg
			t.stdinTail.maxQueueSize = *nLines
			doTail(t, xOpen(t.filename))
		}
	} else {
		t := &stdinTail{maxQueueSize: *nLines}
		doTail(t, os.Stdin)
	}
}
