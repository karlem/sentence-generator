package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/karelm/sentence-generator/generator"
)

func stress() {
	fmt.Println("Running stress test....")
	t := time.Now()
	gen := generator.NewGenerator()
	var wg sync.WaitGroup

	wg.Add(10)
	go func() {
		for index := 0; index < 10; index++ {
			go (func() {
				f, err := os.Open("bible.txt")
				if err != nil {
					panic(err)
				}

				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					t := scanner.Text()

					// No need to try process
					if len(t) == 0 {
						continue
					}

					gen.Learn(t)
				}
				if err := scanner.Err(); err != nil {
					log.Printf("/learn - Error occured during learning: %s", err)
					return
				}

				wg.Done()
			})()
		}
	}()

	txtChan := make(chan string)

	wg.Add(100)
	go func() {
		for index := 0; index < 100; index++ {
			if index > 40 {
				time.Sleep(time.Duration(rand.Intn(600)) * time.Millisecond)
			}

			go (func(gen generator.Generator) {
				txtChan <- gen.Generate()
				wg.Done()

			})(gen)
		}
	}()

	doneChan := make(chan struct{})

	go func() {
		wg.Wait()
		doneChan <- struct{}{}
	}()

	for {
		select {
		case text := <-txtChan:
			fmt.Println(text)
		case <-doneChan:
			fmt.Printf("Test finished after: %s \n", time.Now().Sub(t))
			os.Exit(0)
		}
	}
}
