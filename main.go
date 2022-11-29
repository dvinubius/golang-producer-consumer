package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

func main() {
	totalProduce := 20
	const prodBuffer = 5
	consumed := 0
	produced := 0

	// producer can make up to 10 goods, then must wait for consumers to take some
	goods := make(chan int, prodBuffer)
	var wg sync.WaitGroup

	// PRODUCER - produces 1 / second
	wg.Add(1)
	go func() {
		fmt.Println()
		color.Cyan("= = = Starting Production = = =")
		fmt.Println()
		for p := 1; p <= totalProduce; p++ {
			produced++
			diff := produced - consumed
			msg := fmt.Sprintf("PRODUCED item %v", p)
			color.Cyan(msg)
			bufferLength := prodBuffer - diff
			if bufferLength < 0 {
				bufferLength = 0
			}
			bufferReport := fmt.Sprintf("Buffer: |%v", strings.Repeat("_|", bufferLength))
			if diff >= prodBuffer {
				color.Red(bufferReport)
			} else {
				color.Yellow(bufferReport)
			}
			goods <- p
			time.Sleep(time.Second)
		}
		close(goods)
		fmt.Println()
		color.Cyan("= = = Finished Production = = =")
		fmt.Println()
		wg.Done()
	}()

	// CONSUMER - consumes at random intervals
	wg.Add(1)
	go func() {
		rand.Seed(time.Now().UnixNano())

		for p := range goods {
			consumed++
			color.Green(fmt.Sprintf("        ->        CONSUMED item %v", p))
			delay := rand.Intn(1000) + 800
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
}
