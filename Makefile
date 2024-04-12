all: build

build:
	go build -o server main.go

run: build
	./server

watch:
	air

clean:
	rm -f server

.PHONY: all build run watch clean
