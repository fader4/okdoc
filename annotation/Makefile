gen:
	ragel -Z -G2 -o lexer.go lexer.rl
	ragel -Z -G2 -o extractor.go extractor.rl
	goyacc -p annotation -o parser.go parser.y
	gofmt -w parser.go extractor.go lexer.go
