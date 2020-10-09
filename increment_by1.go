package main

import (
    "fmt";
    //"time";
    "strings";
)

type options struct {
    rollover bool
}

func wearedone() {
    fmt.Println("Done!")
}

func worker(messages chan string) {
    for {
        msg := <-messages
        fmt.Println(msg)
    }
}

func incby1(items []int, messages chan string, opt options) {
    j := 0;
    n := len(items)-1
    for i := n; i > 0; i-- {
        if (j == 0) {
            items[i]++
        }
        if (items[i] > 9) {
            items[i] = 0
            items[i-1]++
            if (items[0] > 9) {
                items[0] = 0
                if (opt.rollover) {
                    items[n]++
                }
            }
        }
        j++
    }
    messages <- "[" + strings.Trim(strings.Replace(fmt.Sprint(items), " ", ", ", -1), "[]") + "]"
	//return items
}

func main() {
    vals := []int{1, 0, 0}

    messages := make(chan string)
    opts := options{rollover: false}

    defer wearedone()

    go worker(messages)

    for n:= 0; n < 900; n++ {
        // fmt.Println(incby1(vals, messages, opts))
        incby1(vals, messages, opts);
    }
}
