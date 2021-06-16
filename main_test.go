// +build unit_test

package main

import (
	"fmt"
	"os"
	"testing"
)

func TestXOpen(t *testing.T) {
	var notWant *os.File
	got := xOpen("testdata/input.txt")
	if got == notWant {
		t.Errorf("not want %p, got %p", notWant, got)
	}
}

func ExampleAppendQueue() {
	stream := xOpen("testdata/input.txt")
	f := &stdinTail{maxQueueSize: defaultNLines}
	f.appendQueue(stream)
	fmt.Println(f.queue)
	// Output:
	// [test090 test091 test092 test093 test094 test095 test096 test097 test098 test99]
}

func ExamplePrintTailQueue() {
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
	filenames := []string{
		"testdata/input.txt",
		"testdata/input2.txt",
	}
	nArg := len(filenames)
	for i := 0; i < len(filenames); i++ {
		stream := xOpen(filenames[i])
		f := &fileTail{
			filename:     filenames[i],
			isNotEndFile: isNotEndFile(i, nArg),
			nArg:         nArg,
			stdinTail: stdinTail{
				maxQueueSize: defaultNLines,
			},
		}
		f.appendQueue(stream)
		f.printTail()
	}
	// Output:
	// ==> testdata/input.txt <==
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
	// ==> testdata/input2.txt <==
	// test097
	// test098
	// test099
}

func ExampleDoTail() {
	stream := xOpen("testdata/input.txt")
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
