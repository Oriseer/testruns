build: 
	@go build -o bin/fs ../testruns/cmd/goapi/main.go

run: build
	./bin/fs

test:
	@go test ./.. -v