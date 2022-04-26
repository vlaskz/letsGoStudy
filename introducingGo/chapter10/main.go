package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		go fmt.Println(n, ":", i)
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("pong %v", i)
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func example() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			Sleeep(1)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			Sleeep(2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second * 3):
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func example2() {
	urls := []string{
		"http://microsoft.com",
		"http://google.com",
		"http://apple.com",
		"http://netflix.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		fmt.Println(result.URL, result.Size)
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println("Biggest HP is:", biggest.URL)

}

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	example()
}

func Sleeep(s int) {

	time.After(time.Second * time.Duration(s))
}
