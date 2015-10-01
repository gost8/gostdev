configure:
	gb vendor update --all

build:
	gofmt -w src/gostdev
	go tool vet src/gostdev/*.go
	gb test
	gb build