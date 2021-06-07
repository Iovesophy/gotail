// +build integration

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

func ExamplePrintTail() {
	f := &stdinTail{
		queue: []string{"test000", "test001", "test002"},
	}
	f.printTail()
	// Output:
	// test000
	// test001
	// test002
}

func ExampleDoTail() {
	stream := xOpen("testdata/test.txt")
	f := &fileTail{
		stdinTail: stdinTail{
			maxQueueSize: defaultNLines,
		},
	}
	doTail(f, stream)
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

func ExampleMain() {
	backup := os.Args
	os.Args = []string{"serial", "-n", "10", "testdata/test.txt"}
	main()
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
	os.Args = backup
}
func ExampleMainMultipleFile() {
	backup := os.Args
	os.Args = []string{"serial", "-n", "10", "testdata/test.txt", "testdata/test.txt"}
	main()
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
	os.Args = backup
}
func ExampleMainStdin() {
	backup := os.Args
	fp, err := os.Open("testdata/test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fp.Close()
	os.Stdin = fp
	os.Args = []string{"serial", "-n", "10"}
	main()
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
	os.Args = backup
}

/* 3 lines */
func ExampleMain3lines() {
	backup := os.Args
	os.Args = []string{"serial", "-n", "10", "testdata/test3lines.txt"}
	main()
	// Output:
	// test097
	// test098
	// test099
	os.Args = backup
}
func ExampleMainMultipleFile3lines() {
	backup := os.Args
	os.Args = []string{"serial", "-n", "10", "testdata/test3lines.txt", "testdata/test3lines.txt"}
	main()
	// Output:
	// ==> testdata/test3lines.txt <==
	// test097
	// test098
	// test099
	//
	// ==> testdata/test3lines.txt <==
	// test097
	// test098
	// test099
	os.Args = backup

	fmt.Println()
}
func ExampleMainStdin3lines() {
	backup := os.Args
	fp, err := os.Open("testdata/test3lines.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fp.Close()
	os.Stdin = fp
	os.Args = []string{"serial", "-n", "10"}
	main()
	// Output:
	// test097
	// test098
	// test099
	os.Args = backup
}
