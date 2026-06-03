package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
	"slices"
	"strings"
)

func readArguments() (int, int, error) {
	flag.Parse() // забрать строку и порезать
	if len(flag.Args()) < 2 {
		fmt.Println("Invalid input. Use 2 arguments") 
		return 0, 0, errors.New("Invalid input")
	}
	var N, err = strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println("Invalid input. First arguments must be integer")
		return 0, 0, errors.New("Invalid input")
	}
	var M, err1 = strconv.Atoi(flag.Args()[1])
	if err1 != nil {
		fmt.Println("Invalid input. Second arguments must be integer")
		return 0, 0, errors.New("Invalid input")
	}
	return N, M, nil
}

func main() {
	var N, M, err = readArguments()
	if err != nil {
		return
	}

	var result []string = make([]string, 0,  N)
	var mu sync.Mutex
	var wg sync.WaitGroup // чтобы мейн подождал горутины

	for i := 0; i < N; i++ {
		wg.Add(1) // добавить 1 
		go func(id int){
			timeSleep := time.Duration(rand.IntN(M))  // сколько наносек, мало
			time.Sleep(timeSleep * time.Millisecond) // сколько наносек в миллисек
			mu.Lock() // если уже заблокирован, то засыпает и в очередь
			result = append(result, fmt.Sprintf("%d, %d", id, timeSleep))
			mu.Unlock()
			wg.Done() // вычесть 
		}(i)
	}
	
	wg.Wait() // ждать мейну

	slices.SortFunc(result, func(i, j string) int {
		a, _ := strconv.Atoi(strings.Fields(i)[1])
		b, _ := strconv.Atoi(strings.Fields(j)[1])
		if  a < b {
			return 1
		} else if a == b{
			return 0
		} else {
			return -1
		}
	})
	for i := 0; i < N; i++ {
		fmt.Println(result[i])
	}
}