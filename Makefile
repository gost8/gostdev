configure:
	gb vendor update --all

build:
	gofmt -w src/gostdev
	gb build