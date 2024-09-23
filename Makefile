gen:
	antlr4 -Dlanguage=Go -o gen antlr/*.g4

.PHONY: gen 
