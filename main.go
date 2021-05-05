package main

import (
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

  c := make(chan string)    // channel where comm is of `string` type

  for _, link := range links {
    go checkLink(link, c)      // `go` is only used before function calls
  }

  fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if err != nil {
    fmt.Println(link, "might be down!")
    c <- "Might be down I think"
    return
  }

  fmt.Println(link, "is up!")
  c <- "Yup, it's up"
}
