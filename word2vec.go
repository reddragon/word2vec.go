package main

import (
  "flag"
  "fmt"
)

func main() {
  var trainFile = flag.String("t", "", "Path of the training file.")
  flag.Parse()
  fmt.Printf("Analyzing training file: %s\n", *trainFile)
}
