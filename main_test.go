// +build integration

package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type dataTests struct {
	all        []string
	defaultArg []string
}

var td = &dataTests{
	all: []string{
		"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test9", "test10",
		"test11", "test12", "test13", "test14", "test15", "test16", "test17", "test18", "test19",
		"test20", "test21", "test22", "test23", "test24", "test25", "test26", "test27", "test28",
		"test29", "test30", "test31", "test32", "test33", "test34", "test35", "test36", "test37",
		"test38", "test39", "test40", "test41", "test42", "test43", "test44", "test45", "test46",
		"test47", "test48", "test49", "test50", "test51", "test52", "test53", "test54", "test55",
		"test56", "test57", "test58", "test59", "test60", "test61", "test62", "test63", "test64",
		"test65", "test66", "test67", "test68", "test69", "test70", "test71", "test72", "test73",
		"test74", "test75", "test76", "test77", "test78", "test79", "test80", "test81", "test82",
		"test83", "test84", "test85", "test86", "test87", "test88", "test89", "test90", "test91",
		"test92", "test93", "test94", "test95", "test96", "test97", "test98", "test99", "test100",
	},
	defaultArg: []string{
		"test91", "test92", "test93", "test94", "test95", "test96", "test97", "test98", "test99", "test100",
	},
}
var count int = len(td.all)

func TestPrintTail(t *testing.T) {
	ExamplePrintTailFile()
	ExamplePrintTailStdin()
}

func ExamplePrintTailFile() {
	stream, err := os.Open("testdata/test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer stream.Close()
	tf := &fileTail{
		nArg:         1,
		filename:     "testdata/test.txt",
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
	tf := &fileTail{
		nArg:         1,
		filename:     "testdata/test.txt",
		isNotEndFile: false,
		stdinTail: stdinTail{
			maxQueueSize: 10,
		},
	}
	stream, err := os.Open(tf.filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer stream.Close()
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
	var got []string
	var want []string
	for i, j := count, 1; i > 0; i, j = i-1, j+1 {
		testAppendQueueForFileTail := &fileTail{
			nArg:         1,
			filename:     "testdata/test.txt",
			isNotEndFile: false,
			stdinTail: stdinTail{
				maxQueueSize: j,
			},
		}
		stream, err := os.Open(testAppendQueueForFileTail.filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer stream.Close()
		testAppendQueueForFileTail.stdinTail.appendQueue(stream)
		got = testAppendQueueForFileTail.stdinTail.queue
		want = td.all[i-1 : count]
		// compare slice by reflect tool
		if flag := reflect.DeepEqual(got, want); flag != true {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}
}

func TestDoTail(t *testing.T) {
	var got []string
	testDoTail := &fileTail{
		nArg:         1,
		filename:     "testdata/test.txt",
		isNotEndFile: false,
		stdinTail: stdinTail{
			maxQueueSize: 10,
		},
	}
	fp, err := os.Open(testDoTail.filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fp.Close()
	doTail(testDoTail, fp)
	got = testDoTail.queue
	for i, want := range td.defaultArg {
		if got[i] != want {
			t.Errorf("got %v\nwant %v", got[i], want)
		}
	}
}

func TestXOpen(t *testing.T) {
	ExampleXOpen()
}

func ExampleXOpen() {
	os.Stdin = xOpen("testdata/test.txt")
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
	os.Args = []string{"serial", "-n", "10", "testdata/test.txt"}

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
	os.Args = []string{"serial", "-n", "10", "testdata/test.txt", "testdata/test.txt"}
	main()
	// Output:
	// ==> testdata/test.txt <==
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
	// ==> testdata/test.txt <==
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
	fp, err := os.Open("testdata/test.txt")
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
