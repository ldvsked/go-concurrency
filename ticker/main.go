package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
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

func ticker(k int, stop chan os.Signal, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		// до сна
		select { // не стопается если канал пуст
		case <-stop:
			fmt.Println("Termination")
			wg.Done()
			return
		default:
			// сколько наносекунд в секунде(число) * количество секунд(но в нано, поэтому надо умножить)
			time.Sleep(time.Second * time.Duration(k))
		}

		// после сна
		select {
		case <-stop:
			fmt.Println("Termination")
			wg.Done()
			return
		default:
			fmt.Printf("Tick %d since %d\n", i+1, (i+1)*k)
		}
	}
}

func main() {
	k, err := readArgumentK()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		fmt.Println("Use 1 uint argument")
		return
	}

	var stop chan os.Signal = make(chan os.Signal, 1) // cap
	// SIGINT(Ctrl+C), SIGTERM - kill
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	var stopTicker chan os.Signal = make(chan os.Signal, 1)
	var val os.Signal

	var wg sync.WaitGroup
	wg.Add(1)
	go ticker(int(k), stopTicker, &wg) // сказано что тиккер не блокирует мейн поэтому горутина

	val = <-stop // ждет пока что-то придет
	stopTicker <- val
	// нужно подождать пока тиккер завершится
	wg.Wait()
}
