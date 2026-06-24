package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func readArgumentK() (uint, error) {
	flag.Parse()

	if len(flag.Args()) != 1 {
		return 0, errors.New("Invalid input. Amount of arguments != 1")
	}

	k, err := strconv.ParseUint(flag.Args()[0], 10, 64)
	if err != nil {
		return 0, errors.New("Invalid input. Argument is not uint")
	}
	return uint(k), nil
}

func ticker(k int) {
	for i := 0; i < k; i++ {
		// сколько наносекунд в секунде(число) * количество секунд(но в нано, поэтому надо умножить)
		time.Sleep(time.Second * time.Duration(k))
		
	}
}

func main() {
	k, err := readArgumentK()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		fmt.Println("Use 1 uint argument")
		return
	}

	

}