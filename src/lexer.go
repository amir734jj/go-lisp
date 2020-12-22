package src

import (
	"fmt"
	"regexp"
)

type Token struct {
	Name    string
	Pattern string
	Value   string
	Ignore  bool
}

type LexerConfig struct {
	Tokens []Token
}

func NewLexer() *LexerConfig {
	l := LexerConfig{
		Tokens: make([]Token, 0),
	}
	return &l
}

func (conf *LexerConfig) Add(token Token) *LexerConfig {
	conf.Tokens = append(conf.Tokens, token)
	return conf
}

func (conf *LexerConfig) Build() func(text string) []Token {
	return func(text string) []Token {
		var buffer string
		tokens := make([]Token, 0)
		for i := 0; i < len(text); i++ {
			buffer = ""
			matches := make([]Token, 0)
			for j := i; j < len(text); j++ {
				buffer = buffer + string(text[j])
				for _, token := range conf.Tokens {
					match, _ := regexp.MatchString(token.Pattern, buffer)
					if match {
						token.Value = buffer
						matches = append(matches, token)
					}
				}
			}

			if len(matches) == 0 {
				_ = fmt.Errorf("failed to find any match: %s", buffer)
				break
			} else {
				var currentMatch = Token{}
				for _, token := range matches {
					if len(token.Value) > len(currentMatch.Value) {
						currentMatch = token
					} else if len(token.Value) == len(currentMatch.Value) {
						for _, t := range conf.Tokens {
							if currentMatch.Name == t.Name {
								break
							} else if token.Name == t.Name {
								currentMatch = Token{Name: t.Name, Value: currentMatch.Value, Pattern: t.Pattern, Ignore: t.Ignore}
								break
							}
						}
					}
				}

				tokens = append(tokens, currentMatch)
				i += len(currentMatch.Value) - 1
			}
		}

		clones := make([]Token, 0)
		for _, token := range tokens {
			if !token.Ignore {
				clones = append(clones, token)
			}
		}

		return clones
	}
}
