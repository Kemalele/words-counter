.PHONY: all
BIN := cmd

all:
	(cd cmd && go build -o ../bin/${BIN} && ../bin/${BIN})