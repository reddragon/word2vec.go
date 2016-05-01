package main

import (
  "bytes"
  "bufio"
  // "errors"
  "flag"
  "fmt"
  // "map"
  "os"
)

type trainParams struct {
  trainPath string
}

func isReadableChar(b byte) bool {
  return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
}

func normalizeChar(b byte) byte {
  if 'A' <= b && b <= 'Z' {
    return 'a' + (b - 'A')
  }
  return b
}

func readFile(params trainParams) {
  file, err := os.Open(params.trainPath)
  if err != nil {
    fmt.Println("Could not open file")
    return
  }
  reader := bufio.NewReader(file)
  var b byte
  var e error

  i := 0
  continuation := false
  var buff bytes.Buffer
  freqCount := make(map[string]int)
  totalWords := 0

  for b, e = reader.ReadByte() ; e == nil; i++ {
    if isReadableChar(b) {
      if !continuation {
        if buff.Len() != 0 {
          str, buffErr := buff.ReadString(' ')
          if buffErr != nil {
            count := freqCount[str] + 1
            freqCount[str] = count
            totalWords = totalWords + 1
            fmt.Printf("Word: %s, Count: %d\n", str, count)
          }
          buff.Reset()
        }
        continuation = true
      }
      buff.WriteByte(normalizeChar(b))
    } else {
      continuation = false
    }
    b, e = reader.ReadByte()
  }
  if buff.Len() != 0 {
    str, buffErr := buff.ReadString(' ')
    if buffErr != nil {
      count := freqCount[str] + 1
      freqCount[str] = count
      totalWords = totalWords + 1
      fmt.Printf("Word: %s, Count: %d\n", str, count)
    }
    buff.Reset()
  }

  fmt.Printf("\n")
  fmt.Println("Error ", e)
  fmt.Printf("Total words = %d, Unique words = %d\n", totalWords, len(freqCount))
}

func main() {
  var params trainParams
  flag.StringVar(&params.trainPath, "t", "", "Path of the training file.")
  flag.Parse()
  fmt.Printf("Analyzing training file: %s\n", params.trainPath)
  readFile(params)
}
