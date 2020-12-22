package src

import "strconv"

func VisitToString(node AST) string {
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
				actualsString += VisitToString(actual)
			} else {
				actualsString += " " + VisitToString(actual)
			}
		}

		return "(" + functionCallToken.Name + actualsString + ")"
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

		return "( defun " + functionDefToken.Name + "(" + formalsString + ") " + VisitToString(functionDefToken.Body) + " )"
	case CondToken:
		condToken := node.(CondToken)
		return "( if " + VisitToString(condToken.Condition) + " " + VisitToString(condToken.If) + " " + VisitToString(condToken.Else) + " )"
	case UnaryOpToken:
		unaryToken := node.(UnaryOpToken)
		return "( " + unaryToken.Op + " " + VisitToString(unaryToken.Expr1) + " )"
	case BinaryOpToken:
		binaryToken := node.(BinaryOpToken)
		return "( " + binaryToken.Op + " " + VisitToString(binaryToken.Expr1) + " " + VisitToString(binaryToken.Expr2) + " )"
	case ParameterToken:
		return node.(ParameterToken).Name
	}

	return ""
}
