package src

import "strconv"

func VisitToCodeGen(node AST) string {
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
				actualsString += VisitToCodeGen(actual)
			} else {
				actualsString += ", " + VisitToCodeGen(actual)
			}
		}

		return functionCallToken.Name + "(" + actualsString + ")"
	case FunctionDefToken:
		functionDefToken := node.(FunctionDefToken)
		formalsString := ""
		for i, formal := range functionDefToken.Formals {
			if i == 0 {
				formalsString += formal
			} else {
				formalsString += ", " + formal
			}
		}

		return "function " + functionDefToken.Name + " (" + formalsString + ") {\n\treturn " + VisitToCodeGen(functionDefToken.Body) + ";\n}"
	case CondToken:
		condToken := node.(CondToken)
		return VisitToCodeGen(condToken.Condition) + " ? " + VisitToCodeGen(condToken.If) + " : " + VisitToCodeGen(condToken.Else)
	case UnaryOpToken:
		unaryToken := node.(UnaryOpToken)
		return "( " + unaryToken.Op + " " + VisitToCodeGen(unaryToken.Expr1) + " )"
	case BinaryOpToken:
		binaryToken := node.(BinaryOpToken)
		return "(" + VisitToCodeGen(binaryToken.Expr1) + " " + binaryToken.Op + " " + VisitToCodeGen(binaryToken.Expr2) + " )"
	case ParameterToken:
		return node.(ParameterToken).Name
	}

	return ""
}

func JavaScriptCodeGen(boundedAsts []BoundedAst) string {
	result := ""
	for _, ast := range boundedAsts {
		result += VisitToCodeGen(ast.AST) + "\n"
	}

	return result
}
