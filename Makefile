gen:
	ragel -Z -G2 -o py_lexer.go py_lexer.rl
	goyacc -p py -o py_parser.go py_parser.y
	gofmt -w py_parser.go py_lexer.go
