package src

import (
	"fmt"
	_ "fmt"
	"strconv"
)

type AST interface{}

type FunctionDefToken struct {
	AST
	Name    string
	Formals []string
	Body    AST
}

type FunctionCallToken struct {
	AST
	Name    string
	Actuals []AST
}

type StringToken struct {
	AST
	Value string
}

type NumberToken struct {
	AST
	Value float64
}

type CondToken struct {
	AST
	Condition AST
	If        AST
	Else      AST
}

type BinaryOpToken struct {
	AST
	Op    string
	Expr1 AST
	Expr2 AST
}

type UnaryOpToken struct {
	AST
	Op    string
	Expr1 AST
	Expr2 AST
}

func ParseNumber(tokens []Token) AST {
	if len(tokens) == 1 && tokens[0].Name == "NUM" {
		f, _ := strconv.ParseFloat(tokens[0].Value, 64)
		return AST(NumberToken{Value: f})
	} else {
		return nil
	}
}

func ParseString(tokens []Token) AST {
	if len(tokens) == 1 && tokens[0].Name == "STRING" {
		return AST(StringToken{Value: tokens[0].Value})
	} else {
		return nil
	}
}

func ParseFunctionDef(tokens []Token) AST {
	if len(tokens) >= 7 && tokens[0].Name == "L_PRN" && tokens[1].Name == "DEFUN" && tokens[len(tokens)-1].Name == "R_PRN" {
		name := tokens[2].Value
		var formals []string
		for i, token := range tokens[4:] {
			if token.Name != "R_PRN" {
				formals = append(formals, token.Value)
			} else {
				tokens = tokens[4+i+1 : len(tokens)-1]
				break
			}
		}
		return FunctionDefToken{Name: name, Formals: formals, Body: Parse(tokens)}
	} else {
		return nil
	}
}

func ParseFunctionCall(tokens []Token) AST {
	if len(tokens) >= 3 && tokens[0].Name == "L_PRN" && tokens[1].Name == "ID" && tokens[len(tokens)-1].Name == "R_PRN" {
		name := tokens[1].Value
		actuals := ParseMultiple(tokens[2 : len(tokens)-1])
		return FunctionCallToken{Name: name, Actuals: actuals}
	} else {
		return nil
	}
}

func ParseCond(tokens []Token) AST {
	if len(tokens) >= 6 && tokens[0].Name == "L_PRN" && tokens[1].Name == "COND" && tokens[len(tokens)-1].Name == "R_PRN" {
		exprs := ParseMultiple(tokens[2 : len(tokens)-1])
		if len(exprs) == 3 {
			return CondToken{Condition: exprs[0], If: exprs[1], Else: exprs[2]}
		}
	}

	return nil
}

func ParseOp(tokens []Token) AST {
	if len(tokens) >= 4 && tokens[0].Name == "L_PRN" && tokens[1].Name == "OP" && tokens[len(tokens)-1].Name == "R_PRN" {
		exprs := ParseMultiple(tokens[2 : len(tokens)-1])
		if len(exprs) == 1 {
			return UnaryOpToken{Op: tokens[1].Value, Expr1: exprs[0]}
		} else if len(exprs) == 2 {
			return BinaryOpToken{Op: tokens[1].Value, Expr1: exprs[0], Expr2: exprs[1]}
		}
	}

	return nil
}

func ParseAtomic(tokens []Token) AST {
	stringToken := ParseString(tokens)
	numberToken := ParseNumber(tokens)

	if stringToken != nil {
		return stringToken
	} else if numberToken != nil {
		return numberToken
	} else {
		return nil
	}
}

func ParseExpression(tokens []Token) AST {
	condToken := ParseCond(tokens)
	functionDefToken := ParseFunctionDef(tokens)
	functionCallToken := ParseFunctionCall(tokens)
	opToken := ParseOp(tokens)

	if condToken != nil {
		return condToken
	} else if functionDefToken != nil {
		return functionDefToken
	} else if functionCallToken != nil {
		return functionCallToken
	} else if opToken != nil {
		return opToken
	} else {
		return nil
	}
}

func Parse(tokens []Token) AST {
	atomicToken := ParseAtomic(tokens)
	exprToken := ParseExpression(tokens)

	if atomicToken != nil {
		return atomicToken
	} else if exprToken != nil {
		return exprToken
	} else {
		return nil
	}
}

func ParseMultiple(tokens []Token) []AST {
	i := 0
	j := i + 1
	var result []AST
	for len(tokens) != 0 && j < len(tokens) {
		temp := Parse(tokens[i:j])
		if temp != nil {
			result = append(result, temp)
			tokens = tokens[j:]
			i = j
			j = i + 1
		} else {
			j++
		}
	}

	if result == nil {
		fmt.Errorf("failed to find any AST: %s", tokens)
	}

	return result
}
