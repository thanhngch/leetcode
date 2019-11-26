package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"testing"
	"time"
)

type Case struct {
	in   string
	want bool
}

func setupData() []Case {
	data, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	dataArray := strings.Split(string(data), "\n")

	cases := []Case{}

	for _, line := range dataArray {
		lineSplit := strings.Split(line, " ")
		cases = append(cases, Case{
			in:   lineSplit[0],
			want: lineSplit[1][0] == '1',
		})
	}
	return cases
}
func TestValidString(t *testing.T) {
	cases := setupData()
	t0 := time.Now()
	for _, c := range cases {
		got := isValid(c.in)
		if got != c.want {
			t.Errorf("isValid(%v) = %v, want %v", c.in, got, c.want)
		} else {
			// fmt.Printf("Test success %s, %v\n", c.in, got)
		}
	}
	t1 := time.Now()
	fmt.Printf("Finish test in %v to run.\n", t1.Sub(t0))
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = Heap = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = TotalHeap = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// func BenchmarkValidString(b *testing.B) {
// 	cases := setupData()
// 	t0 := time.Now()
// 	for i := 0; i < 1000; i++ {
// 		for _, c := range cases {
// 			got := isValid(c.in)
// 			if got != c.want {
// 				b.Errorf("isValid(%v) = %v, want %v", c.in, got, c.want)
// 			}
// 		}
// 	}
// 	t1 := time.Now()
// 	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
// }
