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

  /*
    Infinite loop that watches the channel c;
    when a value comes out of it (from the checkLink function),
    assign its value to the link `l`.

    Given that the `c` channel produces a string, Golang accepts that
    checkLink is called with a link `l` to the channel `c` that itself
    produces a string

    Using a "function literal" (equivalent to lambda in C++ and Python)
    so that the main routine isn't blocked by the `time-Sleep` function
    Here the link is passed as a value so that it's copied to guarantee
    that the child routine is using the same as it had when it was first
    called. Otherwise it would use the value of l which would always be
    changing as the for-loop changed it.
  */
  for l := range c {
    go func(link string) {
        time.Sleep(5 * time.Second)
        checkLink(link, c)
    }(l)
  }
}

// The link with which the function is called gets returned
// through the c channel to be used again in the infinite loop
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
