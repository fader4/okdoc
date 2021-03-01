gen:
	ragel -Z -G2 -o starlark_lexer.go starlark_lexer.rl
	goyacc -p starlark -o starlark_parser.go starlark_parser.y
	gofmt -w starlark_parser.go starlark_lexer.go
