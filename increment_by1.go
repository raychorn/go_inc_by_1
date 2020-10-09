package main

import (
    "fmt"
)

type options struct {
    rollover bool
}

func incby1(items []int, opt options) []int {
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
	return items
}

func main() {
    vals := []int{1, 0, 0}

    opts := options{rollover: false}

    for n:= 0; n < 900; n++ {
	    fmt.Println(incby1(vals, opts))
    }
}
