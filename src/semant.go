package src

type Contour struct {
	table  map[string]BoundedAst
	parent *Contour
}

func newContour() Contour {
	return Contour{table: map[string]BoundedAst{}, parent: nil}
}

func (contour *Contour) searchContour(key string) *BoundedAst {
	if contour == nil {
		println("Unbound variable: " + key)
		return nil
	} else if val, ok := contour.table[key]; ok {
		return &val
	} else {
		return contour.parent.searchContour(key)
	}
}

func (contour *Contour) pushContour() Contour {
	return Contour{table: map[string]BoundedAst{}, parent: contour}
}

func (contour *Contour) popContour() Contour {
	return *contour.parent
}

type BoundedAst struct {
	AST
	Binding *BoundedAst
}

func VisitSemant(node AST, contour *Contour) *BoundedAst {
	switch node.(type) {
	case StringToken:
		return &BoundedAst{AST: node}
	case NumberToken:
		return &BoundedAst{AST: node}
	case FunctionCallToken:
		functionCallToken := node.(FunctionCallToken)
		functionDefBoundedAst := contour.searchContour(functionCallToken.Name)
		functionCallBondedAst := BoundedAst{AST: functionCallToken, Binding: functionDefBoundedAst}

		if functionDefBoundedAst.Binding != nil && len(functionDefBoundedAst.Binding.AST.(FunctionDefToken).Formals) != len(functionCallToken.Actuals) {
			println("Count of actuals and formals are different")
		}

		for _, actual := range functionCallToken.Actuals {
			VisitSemant(actual, contour)
		}

		return &functionCallBondedAst
	case FunctionDefToken:
		functionDefToken := node.(FunctionDefToken)
		functionDefBoundedAst := BoundedAst{AST: functionDefToken}
		contour.table[functionDefToken.Name] = functionDefBoundedAst

		tempContour := contour.pushContour()
		for _, formal := range functionDefToken.Formals {
			tempContour.table[formal] = BoundedAst{AST: ParameterToken{Name: formal}, Binding: &functionDefBoundedAst}
		}

		VisitSemant(functionDefToken.Body, &tempContour)

		tempContour = tempContour.popContour()

		return &functionDefBoundedAst
	case CondToken:
		condToken := node.(CondToken)
		condBoundedAst := BoundedAst{AST: condToken}
		VisitSemant(condToken.Condition, contour)
		VisitSemant(condToken.If, contour)
		VisitSemant(condToken.Else, contour)
		return &condBoundedAst
	case UnaryOpToken:
		unaryToken := node.(UnaryOpToken)
		unaryBoundedAst := BoundedAst{AST: unaryToken}
		return &unaryBoundedAst
	case BinaryOpToken:
		binaryToken := node.(BinaryOpToken)
		binaryBoundedAst := BoundedAst{AST: binaryToken}
		return &binaryBoundedAst
	case ParameterToken:
		parameterToken := node.(ParameterToken)
		return contour.searchContour(parameterToken.Name)
	}

	return nil
}

func SemantAnalyze(asts []AST) []BoundedAst {
	rootContour := newContour()

	var result []BoundedAst
	for _, ast := range asts {
		result = append(result, *VisitSemant(ast, &rootContour))
	}

	return result
}
