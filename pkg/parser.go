package pkg

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

func Parser(f1 [][]string, f2 [][]string) {
	var i = 0.0
	var endFile = make([]string, 0)

	wg := &sync.WaitGroup{}

	for _, f1 := range f1 {
		wg.Add(1)
		ProcentBar(i)
		go parseGo(f1, f2, &endFile, wg)
		i++
		wg.Wait()
	}
	defer FileSave(endFile)
}

func parseGo(f1 []string, f2 [][]string, endFile *[]string, wg *sync.WaitGroup) {
	var b = false
	for _, f2r := range f2 {
		if strings.Contains(f1[0], f2r[0]) {
			b = false
			break
		} else {
			b = true
		}

	}
	if b == true {
		*endFile = append(*endFile, f1[0])
		fmt.Println("Added - ", f1[0])
	}
	wg.Done()
}

/*
Test ProcentBar
*/
func ProcentBar(i float64) {
	var b int
	b = int(math.Ceil(i / 15.0 * 10.0))
	time.Sleep(1 * time.Millisecond)
	h := strings.Repeat("|", b+1) + strings.Repeat("-", 10-b)
	fmt.Printf("\r[%s] %d%% ", h, b*10)
}
