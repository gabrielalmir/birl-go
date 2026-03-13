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
	"IMC": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("IMC ESPERAVA 2 ARGUMENTOS (PESO, ALTURA), FRANGO. RECEBEU %d", len(args))
			}

			var peso, altura float64

			// Helper para extrair número de Int ou Float
			extract := func(obj object.Object) (float64, bool) {
				switch v := obj.(type) {
				case *object.Integer:
					return float64(v.Value), true
				case *object.Float:
					return v.Value, true
				case *object.String:
					if f, err := strconv.ParseFloat(v.Value, 64); err == nil {
						return f, true
					}
				}
				return 0, false
			}

			var ok bool
			if peso, ok = extract(args[0]); !ok {
				return newError("PESO INVÁLIDO OU NÃO NUMÉRICO, TIPO %s", args[0].Type())
			}
			if altura, ok = extract(args[1]); !ok {
				return newError("ALTURA INVÁLIDA OU NÃO NUMÉRICA, TIPO %s", args[1].Type())
			}

			if altura == 0 {
				return newError("DIVISÃO POR ZERO? TA QUERENDO QUEBRAR A JAULA?")
			}

			resultado := peso / (altura * altura)
			return &object.Float{Value: resultado}
		},
	},
}
