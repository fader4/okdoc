gen:
	ragel -Z -G2 -o lexer.go lexer.rl
	ragel -Z -G2 -o preprocessing.go preprocessing.rl
	goyacc -p annotation -o parser.go parser.y
	gofmt -w parser.go preprocessing.go lexer.go
