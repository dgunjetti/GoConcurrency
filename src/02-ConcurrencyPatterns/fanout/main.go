package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	// START1 OMIT
	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	// END1 OMIT

	// START2 OMIT
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	// END2 OMIT

	// START3 OMIT
	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}
	// END3 OMIT

	// START4 OMIT
	primeFinder :=
		func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
			primeStream := make(chan interface{})
			go func() {
				defer close(primeStream)
				for integer := range intStream {
					integer -= 1
					prime := true
					for divisor := integer - 1; divisor > 1; divisor-- {
						if integer%divisor == 0 {
							prime = false
							break
						}
					}
					// END4 OMIT
					// START5 OMIT
					if prime {
						select {
						case <-done:
							return
						case primeStream <- integer:
						}
					}
				}
			}()
			return primeStream
		}
	// END5 OMIT

	// START6 OMIT
	fanIn := func(
		done <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} { // <1>
		var wg sync.WaitGroup // <2>
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) { // <3>
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}
		// END6 OMIT
		// START7 OMIT
		// Select from all the channels
		wg.Add(len(channels)) // <4>
		for _, c := range channels {
			go multiplex(c)
		}
		// Wait for all the reads to complete
		go func() { // <5>
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}
	// END7 OMIT

	// START8 OMIT
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	rand := func() interface{} { return rand.Intn(50000000) }

	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	// END8 OMIT

	// START9 OMIT
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
	// END9 OMIT
}
