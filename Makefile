.PHONY: init update start

yacc_generate:
	@goyacc -o query.y.go query.y
