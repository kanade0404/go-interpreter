package evaluator

import (
	"interpreter/ast"
	"interpreter/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{
		Value: false,
	}
	NULL = &object.Null{}
)

/*
Eval
評価を行う
*/
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// 文
	case *ast.Program:
		return evalStatements(node.Statements)
	// 式
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// 整数値リテラル
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	// 真偽値
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}
	return nil
}

/*
evalStatements
式の評価を行う
*/
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

/*
nativeBoolToBooleanObject
真偽値リテラルの評価
 */
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

/*
evalBangOperatorExpression
論理否定演算子の評価
 */
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

/**
evalMinusPrefixOperatorExpression
符号反転前置演算子の評価
 */
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}