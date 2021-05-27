package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func ExtractStdout(t *testing.T, fnc func()) string {
	t.Helper()
	orgStdout := os.Stdout
	defer func() {
		os.Stdout = orgStdout
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w
	fnc()
	w.Close()
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	s := buf.String()
	// Pipeを使用した際にbuf末尾に改行が追加される問題の対処のため
	return s[:len(s)-1]
}

const count int = 100

func TestPrintTail(t *testing.T) {
	stream, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
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
	actualAll := ExtractStdout(t, tf.stdinTail.printTail)
	expectedAll := "test91\ntest92\ntest93\ntest94\ntest95\ntest96\ntest97\ntest98\ntest99\ntest100"
	if reflect.DeepEqual(actualAll, expectedAll) {
		t.Log(reflect.DeepEqual(actualAll, expectedAll))
	} else {
		t.Errorf("got %v\nwant %v", actualAll, expectedAll)
	}
	// stdin ver check
	os.Stdin = stream
	tf.appendQueue(os.Stdin)
	actualAll = ExtractStdout(t, tf.printTail)
	expectedAll = "test91\ntest92\ntest93\ntest94\ntest95\ntest96\ntest97\ntest98\ntest99\ntest100"
	if reflect.DeepEqual(actualAll, expectedAll) {
		t.Log(reflect.DeepEqual(actualAll, expectedAll))
	} else {
		t.Errorf("got %v\nwant %v", actualAll, expectedAll)
	}
	tf.printTail()
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
			fmt.Fprintf(os.Stderr, "%s\n", err)
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
	// stdin ver check
	expectedAll = []string{}
	for i, nAll := count, 1; i > 0; i, nAll = i-1, nAll+1 {
		testAppendQueueForStdinTail := &stdinTail{
			maxQueueSize: nAll,
		}
		stream, err := os.Open("test.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		defer stream.Close()
		os.Stdin = stream
		testAppendQueueForStdinTail.appendQueue(os.Stdin)
		actualAll = testAppendQueueForStdinTail.queueData
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
		fmt.Fprintf(os.Stderr, "%s\n", err)
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
	var actualAll []string
	var expectedAll []string
	testXOpen := &fileTail{
		nArg:         1,
		filename:     "test.txt",
		isNotEndFile: false,
		stdinTail: stdinTail{
			maxQueueSize: 10,
		},
	}
	stream := xOpen("test.txt")
	doTail(testXOpen, stream)
	actualAll = testXOpen.queueData
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer fp.Close()
	doTail(testXOpen, fp)
	expectedAll = testXOpen.queueData
	if reflect.DeepEqual(actualAll, expectedAll) {
		t.Log(reflect.DeepEqual(actualAll, expectedAll))
	} else {
		t.Errorf("got %v\nwant %v", actualAll, expectedAll)
	}
}

func TestMain(m *testing.T) {
	m.Run("default", func(m *testing.T) {
		backupArgs := os.Args
		testArgs := []string{"serial", "-n", "10", "./test.txt"}
		os.Args = testArgs
		m.Log(os.Args)
		main()
		actualAll := ExtractStdout(m, main)
		expectedAll := "test91\ntest92\ntest93\ntest94\ntest95\ntest96\ntest97\ntest98\ntest99\ntest100"
		if reflect.DeepEqual(actualAll, expectedAll) {
			m.Log(reflect.DeepEqual(actualAll, expectedAll))
		} else {
			m.Errorf("got %v\nwant %v", actualAll, expectedAll)
		}
		os.Args = backupArgs
	})

	m.Run("files", func(m *testing.T) {
		backupArgs := os.Args
		testArgs := []string{"serial", "-n", "10", "./test.txt", "./test.txt"}
		os.Args = testArgs
		m.Log(os.Args)
		main()
		actualAll := ExtractStdout(m, main)
		expectedAll := "==> ./test.txt <==\ntest91\ntest92\ntest93\ntest94\ntest95\ntest96\ntest97\ntest98\ntest99\ntest100\n\n==> ./test.txt <==\ntest91\ntest92\ntest93\ntest94\ntest95\ntest96\ntest97\ntest98\ntest99\ntest100"
		if reflect.DeepEqual(actualAll, expectedAll) {
			m.Log(reflect.DeepEqual(actualAll, expectedAll))
		} else {
			m.Errorf("got %v\nwant %v", actualAll, expectedAll)
		}
		os.Args = backupArgs
	})

	m.Run("stdin", func(m *testing.T) {
		backupArgs := os.Args
		fp, err := os.Open("./test.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		defer fp.Close()
		os.Stdin = fp
		testArgs := []string{"serial", "-n", "10"}
		os.Args = testArgs
		m.Log(os.Args)
		actualAll := ExtractStdout(m, main)
		expectedAll := "test91\ntest92\ntest93\ntest94\ntest95\ntest96\ntest97\ntest98\ntest99\ntest100"
		if reflect.DeepEqual(actualAll, expectedAll) {
			m.Log(reflect.DeepEqual(actualAll, expectedAll))
		} else {
			m.Errorf("got %v\nwant %v", actualAll, expectedAll)
		}
		os.Args = backupArgs
	})
}
