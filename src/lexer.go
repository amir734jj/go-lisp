package src

import (
	"fmt"
	"regexp"
)

type Token struct {
	Name    string
	Pattern string
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

type Pair struct {
	key, value string
}

func (conf *LexerConfig) Add(token Token) *LexerConfig {
	conf.Tokens = append(conf.Tokens, token)
	return conf
}

func (conf *LexerConfig) Build() func(text string) []Pair {
	return func(text string) []Pair {
		var buffer string
		tokens := make([]Pair, 0)
		for i := 0; i < len(text); i++ {
			buffer = ""
			matches := make(map[string]string)
			for j := i; j < len(text); j++ {
				buffer = buffer + string(text[j])
				for _, token := range conf.Tokens {
					match, _ := regexp.MatchString(token.Pattern, buffer)
					if match {
						matches[token.Name] = buffer
					}
				}
			}

			if len(matches) == 0 {
				_ = fmt.Errorf("failed to find any match: %s", buffer)
				break
			} else {
				var longestMatch = Pair{}
				for k, v := range matches {
					if len(v) > len(longestMatch.value) {
						longestMatch = Pair{k, v}
					}
				}

				tokens = append(tokens, longestMatch)
				i += len(longestMatch.value) - 1
			}
		}

		return tokens
	}
}
