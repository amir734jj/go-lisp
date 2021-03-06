package src

import "strconv"

func VisitToLispString(node AST) string {
	switch node.(type) {
	case StringToken:
		return node.(StringToken).Value
	case NumberToken:
		return strconv.FormatFloat(node.(NumberToken).Value, 'f', 6, 64)
	case FunctionCallToken:
		functionCallToken := node.(FunctionCallToken)
		actualsString := ""
		for i, actual := range functionCallToken.Actuals {
			if i == 0 {
				actualsString += VisitToLispString(actual)
			} else {
				actualsString += " " + VisitToLispString(actual)
			}
		}

		return "(" + functionCallToken.Name + " " + actualsString + ")"
	case FunctionDefToken:
		functionDefToken := node.(FunctionDefToken)
		formalsString := ""
		for i, formal := range functionDefToken.Formals {
			if i == 0 {
				formalsString += formal
			} else {
				formalsString += " " + formal
			}
		}

		return "( defun " + functionDefToken.Name + "(" + formalsString + ") " + VisitToLispString(functionDefToken.Body) + " )"
	case CondToken:
		condToken := node.(CondToken)
		return "( if " + VisitToLispString(condToken.Condition) + " " + VisitToLispString(condToken.If) + " " + VisitToLispString(condToken.Else) + " )"
	case UnaryOpToken:
		unaryToken := node.(UnaryOpToken)
		return "( " + unaryToken.Op + " " + VisitToLispString(unaryToken.Expr1) + " )"
	case BinaryOpToken:
		binaryToken := node.(BinaryOpToken)
		return "( " + binaryToken.Op + " " + VisitToLispString(binaryToken.Expr1) + " " + VisitToLispString(binaryToken.Expr2) + " )"
	case ParameterToken:
		return node.(ParameterToken).Name
	}

	return ""
}

func ToLispString(ast []AST) string {
	result := ""
	for _, ast := range ast {
		result += VisitToLispString(ast) + "\n"
	}

	return result
}
