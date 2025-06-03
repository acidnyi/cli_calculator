package main

import (
  "fmt"
  "os"
)

func main(){
  if len(os.Args) != 2 {
    fmt.Println("Usage: calc \"expresion\"")
    fmt.Println("Example of the expresion: 3 + 2 / 4 * (2 - 1)")
    fmt.Println("Supported operations: \n1.+\n2.-\n3.*\n4./\n5.^")
    return
  }
  
  input := os.Args[1]

  tokens, err := Tokenize(input)

  if err != nil {
    fmt.Println("Tokenizer error:", err)
    return 
  }
  
  parser := newParser(tokens)

  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Error: ", r)
    }
  }()

  result := parser.Parse()
  fmt.Printf("= %.2f\n", result)
}
