package main

import (
    "time"
    "fmt"
    "net/http"
)

func main() {
  links := []string {
    "https://google.com",
    "https://facebook.com",
    "https://amazon.com",
    "https://golang.org",
  }

  c := make(chan string)       // channel where comm is of `string` type

  for _, link := range links {
    go checkLink(link, c)      // `go` is only used before function calls
  }

  for l := range c {
    go func(link string) {
        time.Sleep(5 * time.Second)
        checkLink(link, c)
    }(l)
  }
}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if err != nil {
    fmt.Println(link, "might be down!")
    c <- link
    return
  }

  fmt.Println(link, "is up!")
  c <- link
}
