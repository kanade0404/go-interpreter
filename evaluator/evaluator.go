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
	// 前置演算子
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	// 中置演算子
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
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

/*
evalInfixExpression
中置演算子式の評価
オペランドを先に評価して両方整数でないならNULLを返す
 */
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch  {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	default:
		return NULL
	}
}

/*
evalIntegerInfixExpression
オペランドの両方が整数である中置演算子式の評価
 */
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}