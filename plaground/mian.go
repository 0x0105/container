// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/sbinet/go-python"

	"gopl.io/ch5/links"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	gotr := "foo"
	pystr := python.(gostr)
	str := python.PyString_AsString(pystr)
	fmt.Println("hello [", str, "]")

	json.MarshalIndent()
	worklist := make(chan []string)

	// Start with the command-line arguments.
	var wg sync.WaitGroup
	go func() { worklist <- []string{"https://books.studygolang.com/gopl-zh"} }()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for _, link := range <-worklist {
		if !seen[link] {
			seen[link] = true
			wg.Add(1)
			go func(link string) {
				defer wg.Done()
				worklist <- crawl(link)
			}(link)
		}
	}
	wg.Wait()

	fmt.Println(seen)
}
