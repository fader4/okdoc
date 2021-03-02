gen:
	ragel -Z -G2 -o lexer.go lexer.rl
	goyacc -p meta -o parser.go parser.y
	gofmt -w parser.go lexer.go
