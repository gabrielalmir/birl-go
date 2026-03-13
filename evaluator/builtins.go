package evaluator

import (
	"birl-go/object"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"TAMANHO": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("TAMANHO ESPERAVA 1 ARGUMENTO, FRANGO. RECEBEU %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("ARGUMENTO PRO TAMANHO NÃO SUPORTADO, TIPO %s", args[0].Type())
			}
		},
	},
	"CONVERTE_MONSTRO": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("CONVERTE_MONSTRO ESPERAVA 1 ARGUMENTO, FRANGO. RECEBEU %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				val, err := strconv.ParseInt(arg.Value, 10, 64)
				if err != nil {
					return newError("NÃO DEU PRA CONVERTER ISSO AÍ NÃO: %s", arg.Value)
				}
				return &object.Integer{Value: val}
			case *object.Integer:
				return arg
			default:
				return newError("ARGUMENTO PRO CONVERTE_MONSTRO NÃO SUPORTADO, TIPO %s", args[0].Type())
			}
		},
	},
}
