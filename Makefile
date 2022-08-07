GOPATH := $(shell go env GOPATH)

.DEFAULT_GOAL := run

generate:
	rm -f schema/flatbuf/*.go schema/karmem/*.go
	flatc --go --grpc -o ./schema ./schema/flatbuf/nemesis.fbs
	go run karmem.org/cmd/karmem build --golang -o "schema/karmem" schema/karmem/nemesis.km

bench:
	go test -bench=. -benchmem -run=^$$ ./...

run:
	@go run ./flatbuffers/main.go
	@go run ./karmem/main.go
