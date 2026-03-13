package evaluator

import (
	"io"
	"fmt"
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

func (e *Evaluator) evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = e.Eval(statement, env)

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
			if rt == object.ERROR_OBJ {
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
