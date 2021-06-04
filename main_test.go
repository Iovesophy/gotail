package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
)

const count int = 100

func TestPrintTail(t *testing.T) {
	ExamplePrintTailFile()
	ExamplePrintTailStdin()
}

func ExamplePrintTailFile() {
	stream, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer stream.Close()
	tf := &fileTail{
		nArg:         1,
		filename:     "test.txt",
		isNotEndFile: false,
		stdinTail: stdinTail{
			maxQueueSize: 10,
		},
	}
	tf.appendQueue(stream)
	tf.stdinTail.printTail()
	// Output:
	// test91
	// test92
	// test93
	// test94
	// test95
	// test96
	// test97
	// test98
	// test99
	// test100
}

func ExamplePrintTailStdin() {
	stream, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer stream.Close()
	tf := &fileTail{
		nArg:         1,
		filename:     "test.txt",
		isNotEndFile: false,
		stdinTail: stdinTail{
			maxQueueSize: 10,
		},
	}
	os.Stdin = stream
	tf.appendQueue(os.Stdin)
	tf.printTail()
	// Output:
	// test91
	// test92
	// test93
	// test94
	// test95
	// test96
	// test97
	// test98
	// test99
	// test100
}

func TestAppendQueue(t *testing.T) {
	var actualAll []string
	var expectedAll []string

	for i, nAll := count, 1; i > 0; i, nAll = i-1, nAll+1 {
		testAppendQueueForFileTail := &fileTail{
			nArg:         1,
			filename:     "test.txt",
			isNotEndFile: false,
			stdinTail: stdinTail{
				maxQueueSize: nAll,
			},
		}
		stream, err := os.Open(testAppendQueueForFileTail.filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer stream.Close()
		testAppendQueueForFileTail.stdinTail.appendQueue(stream)
		actualAll = testAppendQueueForFileTail.stdinTail.queueData
		expectedAll = append([]string{"test" + strconv.Itoa(i)}, expectedAll...)
		if reflect.DeepEqual(actualAll, expectedAll) {
			t.Log(reflect.DeepEqual(actualAll, expectedAll))
		} else {
			t.Errorf("got %v\nwant %v", actualAll, expectedAll)
		}
	}
}

func TestDoTail(t *testing.T) {
	var actualAll []string
	var expectedAll []string

	testDoTail := &fileTail{
		nArg:         1,
		filename:     "test.txt",
		isNotEndFile: false,
		stdinTail: stdinTail{
			maxQueueSize: 10,
		},
	}
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fp.Close()
	doTail(testDoTail, fp)
	actualAll = testDoTail.queueData
	for i, nAll := count, 1; i > count-10; i, nAll = i-1, nAll+1 {
		expectedAll = append([]string{"test" + strconv.Itoa(i)}, expectedAll...)
	}
	if reflect.DeepEqual(actualAll, expectedAll) {
		t.Log(reflect.DeepEqual(actualAll, expectedAll))
	} else {
		t.Errorf("got %v\nwant %v", actualAll, expectedAll)
	}
}

func TestXOpen(t *testing.T) {
	ExampleXOpen()
}

func ExampleXOpen() {
	os.Stdin = xOpen("test.txt")
	backup := os.Args
	os.Args = []string{"serial", "-n", "10"}
	main()
	// Output:
	// test91
	// test92
	// test93
	// test94
	// test95
	// test96
	// test97
	// test98
	// test99
	// test100
	os.Args = backup
}

func TestMain(t *testing.T) {
	ExampleMainDefault()
	ExampleMainMultipleFile()
	ExampleMainStdin()
}

func ExampleMainDefault() {
	backup := os.Args
	os.Args = []string{"serial", "-n", "10", "test.txt"}

	main()
	// Output:
	// test91
	// test92
	// test93
	// test94
	// test95
	// test96
	// test97
	// test98
	// test99
	// test100
	os.Args = backup
}

func ExampleMainMultipleFile() {
	backup := os.Args
	os.Args = []string{"serial", "-n", "10", "test.txt", "test.txt"}
	main()
	// Output:
	// ==> test.txt <==
	// test91
	// test92
	// test93
	// test94
	// test95
	// test96
	// test97
	// test98
	// test99
	// test100
	//
	// ==> test.txt <==
	// test91
	// test92
	// test93
	// test94
	// test95
	// test96
	// test97
	// test98
	// test99
	// test100
	os.Args = backup
}

func ExampleMainStdin() {
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fp.Close()
	os.Stdin = fp
	backup := os.Args
	os.Args = []string{"serial", "-n", "10"}
	main()
	// Output:
	// test91
	// test92
	// test93
	// test94
	// test95
	// test96
	// test97
	// test98
	// test99
	// test100
	os.Args = backup
}
