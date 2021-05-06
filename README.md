# Channels Exercise in Golan

The code is for the Go-Channels project from the Udemy course [Go: The Complete Deeloper's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)

The application checks regularly, on a 5 seconds interval, the state of the websites listed in the `links` slice of strings.

To avoid a potential hang-up waiting for a page to respond when the application would loop throught the slice of strings, checking one webpage at a time, [Goroutines](https://tour.golang.org/concurrency/1) are implemented.

## Functionality

- `func main()`
Loops through the `links` slice of strings once, creating a `goroutine` that will handle the `checkLink(link, c)` function. In its call, the `channel` `c` is given as an input.

After this initial loop, a second loop will watch the `channel` `c` (indefinitely). When a value comes out of it (from the `checkLink` function called previously), the value is assigned to a `link` `l`. This link is then used in a [function literal](https://golang.org/ref/spec#Function_literals) (equivalent to a lambda in C++ or Python). The function literal is used so that the main routine isn't blocked by the `time.Sleep()` function.
The `link` `l` is passed as a value

```bash
    go func(link string) {
    [...]
    } (l)
```

so that it's copied to guarantee that the child routine is using the same value that `l` pointed to when the child routine was created, instead of whichever one `l` was pointing at in the main for-loop.

- `func checkLink(link string, c chan string)`

For every call it checks the state for the webpage indicated by `link` and verifies that there is no `error` `err` returned. If there is (f. ex. the page is down) it prints "might be down!" and exits the function.

If there is no error returned by the `http.Get(link)` call, it indicates that the link is up and the `link` with which the function is called gets returned through `channel` `c` to be used again in the infinite loop.

## Output

A continuous list of strings indicating if the webpages in the `links` slice of strings are up / down, until there is an interrupt signal from the keyboard (Ctrl+C).

### Example

```bash
https://golang.org is up!
https://google.com is up!
https://facebook.com is up!
https://amazon.com is up!
https://golang.org is up!
https://google.com is up!
https://facebook.com is up!
```

## How-To

### Build and Run the Application

```bash
go build main.go
./main
```

### Test

no tests were developed for this application

### Notes

N/A
