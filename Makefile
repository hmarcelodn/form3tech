build:
	go run main.go

tests:
	cd ./test && go test *.go -v -cover
	