package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught: ", sig)
}

func main() {
	fmt.Printf("Process IO: %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)
	start := time.Now()

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execution time: ", duration)
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(0)
			default:
				fmt.Println("Caught: ", sig)
			}
		}
	}()
	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}

}
