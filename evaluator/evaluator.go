package evaluator

import (
	"io"
	"fmt"
	"strconv"
	"birl-go/ast"
	"birl-go/object"
)

var (
	NULL = &object.Null{}
)

type Evaluator struct {
	Stdout io.Writer
}

func New(out io.Writer) *Evaluator {
	return &Evaluator{Stdout: out}
}

func (e *Evaluator) Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return e.evalProgram(node, env)

	case *ast.BlockStatement:
		return e.evalBlockStatement(node, env)

	case *ast.ExpressionStatement:
		return e.Eval(node.Expression, env)

	case *ast.VarStatement:
		val := e.Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return NULL

	case *ast.PrintStatement:
		val := e.Eval(node.Expression, env)
		if isError(val) {
			return val
		}
		fmt.Fprintf(e.Stdout, "%s\n", val.Inspect())
		return NULL

	case *ast.ReadStatement:
		var input string
		fmt.Scan(&input)
		// Tenta converter para int se possível, senão guarda como string
		if val, err := strconv.ParseInt(input, 10, 64); err == nil {
			obj := &object.Integer{Value: val}
			if ident, ok := node.Expression.(*ast.Identifier); ok {
				env.Set(ident.Value, obj)
			}
			return obj
		}
		obj := &object.String{Value: input}
		if ident, ok := node.Expression.(*ast.Identifier); ok {
			env.Set(ident.Value, obj)
		}
		return obj

	case *ast.IfStatement:
		return e.evalIfStatement(node, env)

	case *ast.WhileStatement:
		return e.evalWhileStatement(node, env)

	case *ast.ReturnStatement:
		val := e.Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.Identifier:
		return e.evalIdentifier(node, env)

	case *ast.InfixExpression:
		left := e.Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := e.Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return e.evalInfixExpression(node.Operator, left, right)
	}

	return nil
}

func (e *Evaluator) evalIfStatement(is *ast.IfStatement, env *object.Environment) object.Object {
	condition := e.Eval(is.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return e.Eval(is.Consequence, env)
	} else if is.Alternative != nil {
		return e.Eval(is.Alternative, env)
	} else {
		return NULL
	}
}

func (e *Evaluator) evalWhileStatement(ws *ast.WhileStatement, env *object.Environment) object.Object {
	var result object.Object = NULL

	for {
		condition := e.Eval(ws.Condition, env)
		if isError(condition) {
			return condition
		}

		if !isTruthy(condition) {
			break
		}

		result = e.Eval(ws.Body, env)
		if result != nil && (result.Type() == object.RETURN_VALUE_OBJ || result.Type() == object.ERROR_OBJ) {
			return result
		}
	}

	return result
}

func isTruthy(obj object.Object) bool {
	switch obj.Type() {
	case object.NULL_OBJ:
		return false
	case object.BOOLEAN_OBJ:
		return obj.(*object.Boolean).Value
	case object.INTEGER_OBJ:
		return obj.(*object.Integer).Value != 0
	default:
		return true
	}
}

func (e *Evaluator) evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = e.Eval(statement, env)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
		if err, ok := result.(*object.Error); ok {
			return err
		}
	}

	return result
}

func (e *Evaluator) evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = e.Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func (e *Evaluator) evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return newError("identificador não encontrado: %s", node.Value)
}

func (e *Evaluator) evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return e.evalIntegerInfixExpression(operator, left, right)
	case operator == "+":
		return &object.String{Value: left.Inspect() + right.Inspect()}
	case operator == "==":
		return &object.Boolean{Value: left.Inspect() == right.Inspect()}
	case operator == "!=":
		return &object.Boolean{Value: left.Inspect() != right.Inspect()}
	default:
		return newError("tipo desconhecido: %s %s %s", left.Type(), operator, right.Type())
	}
}

func (e *Evaluator) evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
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
		return &object.Boolean{Value: leftVal < rightVal}
	case ">":
		return &object.Boolean{Value: leftVal > rightVal}
	case "==":
		return &object.Boolean{Value: leftVal == rightVal}
	case "!=":
		return &object.Boolean{Value: leftVal != rightVal}
	default:
		return newError("operador desconhecido: %s", operator)
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
