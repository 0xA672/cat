package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    if len(os.Args) == 1 {
        io.Copy(os.Stdout, os.Stdin)
        return
    }

    for _, file := range os.Args[1:] {
        f, err := os.Open(file)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cat: %v\n", err)
            os.Exit(1)
        }
        io.Copy(os.Stdout, f)
        f.Close()
    }
}
