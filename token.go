package main

import (
  "fmt"
  "unicode"
)

type TokenType int

const (
  EOF TokenType = iota
  NUMBER
  PLUS
  MINUS
  MUL
  DIV
  LPAREN
  RPAREN
  POW
  IDENTIFIER
)

type Token struct {
  Type TokenType
  Value string
}

func Tokenize(input string) ([]Token, error) {
  var tokens []Token

  i := 0
  for i < len(input){
    ch := input[i]

    if unicode.IsSpace(rune(ch)) {
      i++
      continue
    }

    switch ch {
    case '+':
      tokens = append(tokens, Token{Type: PLUS, Value: "+"})
    case '-':
      tokens = append(tokens, Token{Type: MINUS, Value: "-"})
    case '*':
      tokens = append(tokens, Token{Type: MUL, Value: "*"})
    case '/':
      tokens = append(tokens, Token{Type: DIV, Value: "/"})
    case '(':
      tokens = append(tokens, Token{Type: LPAREN, Value: "("})
    case ')':
      tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
    case '^':
      tokens = append(tokens, Token{Type: POW, Value: "^"})
    default:
      if unicode.IsDigit(rune(ch)) || ch == '.' {
        start := i

        for i < len(input) && (unicode.IsDigit(rune(input[i])) || input[i] == '.') {
          i++
        }

        tokens = append(tokens, Token{Type: NUMBER, Value: input[start:i]})
        continue
      } else if unicode.IsLetter(rune(ch)) {
        
        start := i
        for i < len(input) && unicode.IsLetter(rune(input[i])){
          i++
        }
        
        tokens = append(tokens, Token{Type: IDENTIFIER, Value: input[start:i]})
        continue
      } else{
        return nil, fmt.Errorf("Unexpected character: %c", ch)
      }

    }

    i++
  }
  tokens = append(tokens, Token{Type: EOF})
  return tokens, nil
}
