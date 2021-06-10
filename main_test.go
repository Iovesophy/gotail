// +build unit_test

package main

import (
	"fmt"
	"os"
	"testing"
)

func TestXOpen(t *testing.T) {
	var notWant *os.File
	got := xOpen("testdata/test.txt")
	if got == notWant {
		t.Errorf("not want %p, got %p", notWant, got)
	}
}

func ExampleAppendQueue() {
	stream := xOpen("testdata/test.txt")
	f := &stdinTail{maxQueueSize: defaultNLines}
	f.appendQueue(stream)
	fmt.Println(f.queue)
	// Output:
	// [test090 test091 test092 test093 test094 test095 test096 test097 test098 test099]
}

func ExampleprintTailQueue() {
	f := &stdinTail{
		queue: []string{"test000", "test001", "test002"},
	}
	f.printTail()
	// Output:
	// test000
	// test001
	// test002
}

func ExamplePrintTailMultipleFile() {
	filename := "testdata/test.txt"
	stream := xOpen(filename)
	f := &fileTail{
		filename:     filename,
		isNotEndFile: true,
		nArg:         2,
		stdinTail: stdinTail{
			maxQueueSize: defaultNLines,
		},
	}
	f.appendQueue(stream)
	f.printTail()
	filename = "testdata/test3lines.txt"
	stream = xOpen(filename)
	f = &fileTail{
		filename:     filename,
		isNotEndFile: false,
		nArg:         2,
		stdinTail: stdinTail{
			maxQueueSize: defaultNLines,
		},
	}
	f.appendQueue(stream)
	f.printTail()
	// Output:
	// ==> testdata/test.txt <==
	// test090
	// test091
	// test092
	// test093
	// test094
	// test095
	// test096
	// test097
	// test098
	// test099
	//
	// ==> testdata/test3lines.txt <==
	// test097
	// test098
	// test099
}

func ExampleDoTail() {
	stream := xOpen("testdata/test.txt")
	f := &fileTail{
		stdinTail: stdinTail{
			maxQueueSize: defaultNLines,
		},
	}
	doTail(f, stream)
	// Output:
	// test090
	// test091
	// test092
	// test093
	// test094
	// test095
	// test096
	// test097
	// test098
	// test099
}

func ExampleParseCLine() {
	backupArgs := os.Args
	os.Args = []string{"serial", "-n", "10", "testdata/test.txt"}
	nFlags, nLines, nArg := parseCLine()
	fmt.Println(nArg, *nLines, nFlags.Arg(0))
	// Output:
	// 1 10 testdata/test.txt
	os.Args = backupArgs
}

func ExampleRecExecFileTail() {
	backupArgs := os.Args
	os.Args = []string{"serial", "-n", "10", "testdata/test.txt"}
	nFlags, nLines, nArg := parseCLine()
	recExec(nFlags, nLines, nArg)
	// Output:
	// test090
	// test091
	// test092
	// test093
	// test094
	// test095
	// test096
	// test097
	// test098
	// test099
	os.Args = backupArgs
}

func ExampleRecExecStdinTail() {
	args, stdin := os.Args, os.Stdin
	os.Args = []string{"serial", "-n", "10"}
	os.Stdin = xOpen("testdata/test.txt")
	nFlags, nLines, nArg := parseCLine()
	recExec(nFlags, nLines, nArg)
	// Output:
	// test090
	// test091
	// test092
	// test093
	// test094
	// test095
	// test096
	// test097
	// test098
	// test099
	os.Args, os.Stdin = args, stdin
}
