package main

import "fmt"

func main() {
    seq := []string{"cat", "cat", "dog", "cat", "tree"}
    set := make(map[string]struct{})

    for _, s := range seq {
        set[s] = struct{}{}
    }

    fmt.Println(set)
}
