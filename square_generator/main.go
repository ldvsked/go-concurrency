package main

import (
	"flag"
	"errors"
	"fmt"
	"strconv"
)

func readArguments() (int, int, error) {
	flag.Parse() // забрать все что передано в командной строке
	if len(flag.Args()) != 2 { // все элементы которые не флаги
		return 0, 0, errors.New("Invalid input")
	}
	var k, errK = strconv.Atoi(flag.Args()[0])
	var n, errN = strconv.Atoi(flag.Args()[1])
	if errK != nil || errN != nil {
		return 0, 0, errors.New("Invalid input")
	}
	return k, n, nil
}

func generator(k, n int) <-chan int {
	var c1 chan int = make(chan int)
	go func() {
		for i := k; i <= n; i++ {
			c1 <- i
		}
		close(c1) // кто создал тот и закрывает, и чтобы range не ждал вечность 
	}()
	return c1
}

func square(c1 <-chan int) <-chan int {
	var c2 chan int = make(chan int) 
	go func() {
		for val := range c1 { // будет считывать пока канал не будет закрыть 
			c2 <- val * val
		}
		close(c2)
	}()
	return c2
}


func main() {
	var k, n, err = readArguments()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Use 2 integer arguments")
		return
	}

	c1 := generator(k, n)
	c2 := square(c1)
	for val := range c2 {
		fmt.Println(val)
	}
}