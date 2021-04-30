package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestTail(t *testing.T) {
	var actual_all []string
	var expected_all []string
	count := 0
	fp, err := os.Open("./test.txt")
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		count++
	}
	n_all := 1
	for i := count; i > 0; i-- {
		fp, err = os.Open("./test.txt")
		actual_all = tail(fp, err, n_all)
		expected_all = append([]string{strconv.Itoa(i)}, expected_all...)
		if reflect.DeepEqual(actual_all, expected_all) {
			t.Log(reflect.DeepEqual(actual_all, expected_all))
		} else {
			t.Errorf("got %v\nwant %v", actual_all, expected_all)
		}
		n_all++
	}

	var actual_neg []string
	var expected_neg []string
	n_neg := -1
	for i := count; i > 0; i-- {
		fp, err = os.Open("./test.txt")
		actual_neg = tail(fp, err, n_neg)
		expected_neg = append([]string{strconv.Itoa(i)}, expected_neg...)
		if reflect.DeepEqual(actual_neg, expected_neg) {
			t.Log(reflect.DeepEqual(actual_neg, expected_neg))
		} else {
			t.Errorf("got %v\nwant %v", actual_neg, expected_neg)
		}
		n_neg--
	}

	var actual_default []string
	var expected_default []string
	fp, err = os.Open("./test.txt")
	actual_default = tail(fp, err, 0)
	for i := count; i > count-10; i-- {
		expected_default = append([]string{strconv.Itoa(i)}, expected_default...)
	}
	if reflect.DeepEqual(actual_default, expected_default) {
		t.Log(reflect.DeepEqual(actual_default, expected_default))
	} else {
		t.Errorf("got %v\nwant %v", actual_default, expected_default)
	}
}

func TestInitQueue(t *testing.T) {
	actual_queue, actual_cursor := init_queue()
	expected_queue := []string{}
	expected_cursor := 0

	if reflect.DeepEqual(actual_queue, expected_queue) {
		t.Log(reflect.DeepEqual(actual_queue, expected_queue))
	} else if reflect.DeepEqual(actual_cursor, expected_cursor) {
		t.Log(reflect.DeepEqual(actual_cursor, expected_cursor))
	} else {
		t.Errorf("got %v\nwant %v", actual_queue, expected_queue)
		t.Errorf("got %v\nwant %v", actual_cursor, expected_cursor)
	}
}

func TestShowQueue(t *testing.T) {
	var actual []string
	var expected []string
	var queue []string
	count := 0
	fp, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		count++
	}
	for i := count; i > 0; i-- {
		fp, err = os.Open("./test.txt")
		queue = append([]string{strconv.Itoa(i)}, expected...)
		actual = show_queue(queue, i)
		expected = []string{}
		if reflect.DeepEqual(actual, expected) {
			t.Log(reflect.DeepEqual(actual, expected))
		} else {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

func TestEnqueue(t *testing.T) {
	queue := []string{"test"}
	value := "test"
	actual := enqueue(queue, value)
	expected := []string{"test", "test"}
	if reflect.DeepEqual(actual, expected) {
		t.Log(reflect.DeepEqual(actual, expected))
	} else {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestDequeue(t *testing.T) {
	queue := []string{"test", "test"}
	actual := dequeue(queue)
	expected := []string{"test"}
	if reflect.DeepEqual(actual, expected) {
		t.Log(reflect.DeepEqual(actual, expected))
	} else {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCallTail(t *testing.T) {
	fp, err := os.Open("./test.txt")
	actual := call_tail(fp, err, 10)
	expected := []string{}
	if reflect.DeepEqual(actual, expected) {
		t.Log(reflect.DeepEqual(actual, expected))
	} else {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMain(m *testing.T) {
	m.Run("default", func(m *testing.T) {
		backupArgs := os.Args
		testArgs := []string{"serial", "-n", "10", "./test.txt"}
		os.Args = testArgs
		m.Log(os.Args)
		main()
		os.Args = backupArgs
	})
}
