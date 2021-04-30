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
	if err != nil {
		fmt.Println("Error: No such file or directory")
		os.Exit(1)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		count++
	}
	n_all := 1
	for i := count; i > 0; i-- {
		fp, err = os.Open("./test.txt")
		actual_all = tail(fp, n_all)
		expected_all = append([]string{strconv.Itoa(i)}, expected_all...)
		if reflect.DeepEqual(actual_all, expected_all) {
			t.Log(reflect.DeepEqual(actual_all, expected_all))
		} else {
			t.Errorf("got %v\nwant %v", actual_all, expected_all)
		}
		n_all++
	}
}

func TestMain(m *testing.T) {
	m.Run("default", func(m *testing.T) {
		backupArgs := os.Args
		testArgs := []string{"serial", "-n", "10", "./test.txt", "./test.txt"}
		os.Args = testArgs
		m.Log(os.Args)
		main()
		os.Args = backupArgs
	})
}
