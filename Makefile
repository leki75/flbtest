GOPATH := $(shell go env GOPATH)

generate:
	rm -f schema/flatbuf/*.go
	flatc --go --grpc -o ./schema ./schema/flatbuf/nemesis.fbs

bench:
	go test -bench=. -benchmem -run=^$$ ./...
