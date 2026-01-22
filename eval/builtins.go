package eval

import (
	"fmt"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want 1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to 'len' not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want 1",
					len(args))
			}

			switch arr := args[0].(type) {
			case *object.Array:
				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				}
				return NULL
			default:
				return newError("argument to 'first' must be Array. got %s",
					args[0].Type())
			}
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want 1",
					len(args))
			}

			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to 'last' must be Array. got %s",
					args[0].Type())
			}

			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},

	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got %d, want 1",
					len(args))
			}

			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to 'rest' must be Array. got %s",
					args[0].Type())
			}

			length := len(arr.Elements)
			if length > 0 {
				newArr := make([]object.Object, length-1, length-1)
				copy(newArr, arr.Elements[1:length])
				return &object.Array{Elements: newArr}
			}

			return NULL
		},
	},

	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got %d, want 2",
					len(args))
			}

			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to 'push' must be Array. got %s",
					args[0].Type())
			}

			length := len(arr.Elements)
			newArr := make([]object.Object, length+1, length+1)
			copy(newArr, append(arr.Elements, args[1]))
			return &object.Array{Elements: newArr}
		},
	},

	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}
