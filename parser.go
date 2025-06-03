package main

import (
  "fmt"
  "strconv"
)

type Parser struct {
  tokens []Token
  pos int
}

func newParser(tokens []Token) *Parser {
  return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) current() Token {
  if p.pos >= len(p.tokens){
    return Token{Type: EOF}
  }
  return p.tokens[p.pos]
}

func (p *Parser) eat(expected TokenType) Token {
  tok := p.current()
  if tok.Type == expected {
    p.pos++
    return tok
  }
  panic(fmt.Sprintf("Expected token %v but got %v", expected, tok.Type))
}

func (p *Parser) Parse() float64 {
  return p.parseExpr()
}

func (p *Parser) parseExpr() float64 {
  result := p.parseTerm()

  for {
    tok := p.current()
    if tok.Type == PLUS {
      p.eat(PLUS)
      result += p.parseTerm()
    } else if tok.Type == MINUS {
      p.eat(MINUS)
      result -= p.parseTerm()
    } else {
      break
    }
  }

  return result
}

func (p *Parser) parseTerm() float64 {
  result := p.parseFactor()

  for {
    tok := p.current()

    if tok.Type == MUL {
      p.eat(MUL)
      result *= p.parseFactor()
    } else if tok.Type == DIV {
      p.eat(DIV)
      denom := p.parseFactor()
      if denom == 0 {
        panic("Division by zero.")
      }
      result /= denom
    } else {
      break
    }
  }
  
  return result
}

func (p *Parser) parseFactor() float64 {
  tok := p.current()

  if tok.Type == NUMBER {
    p.eat(NUMBER)
    num, _ := strconv.ParseFloat(tok.Value, 64)
    return num
  } else if tok.Type == LPAREN {
    p.eat(LPAREN)
    result := p.parseExpr()
    p.eat(RPAREN)
    return result
  }

  panic(fmt.Sprintf("Unexpected token: %v", tok))
}
