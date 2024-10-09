package main

import (
    "errors"
    "fmt"
)

var (
    ErrChunkNotFound = errors.New("chunk not found")
)

func HandleError(err error) {
    if errors.Is(err, ErrChunkNotFound) {
        fmt.Println("Handle chunk not found error")
    } else {
        fmt.Println("Handle general error")
    }
}
