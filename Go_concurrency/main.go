package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func sendRequest(url string) {
	defer wg.Done()
	res, err := http.Get(url)

	mut.Lock()
	defer mut.Unlock()
	fmt.Printf("[%d] %s\n", res.StatusCode, url)

	if err == nil {

	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage : go run main.og <url1>")
	}

	for _, url := range os.Args[1:] {
		go sendRequest("https://" + url)
		wg.Add(1)
	}

	wg.Wait()
}
