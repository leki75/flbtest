GOPATH := $(shell go env GOPATH)

.DEFAULT_GOAL := run

generate:
	rm -f schema/flatbuf/*.go schema/karmem/*.go schema/proto/*.go
		protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-vtproto_out=. --plugin protoc-gen-go-vtproto="$(GOPATH)/bin/protoc-gen-go-vtproto" \
		--go-vtproto_opt=paths=source_relative \
		--go-vtproto_opt=features=marshal+unmarshal+size \
		./schema/proto/nemesis.proto
	flatc --go -o ./schema ./schema/flatbuf/nemesis.fbs
	go run karmem.org/cmd/karmem build --golang -o "schema/karmem" schema/karmem/nemesis.km

bench:
	go test -bench=. -benchmem -run=^$$ ./...

run:
	@go run cmd/proto/main.go
	@go run cmd/flatbuffers/main.go
	@go run cmd/karmem/main.go
