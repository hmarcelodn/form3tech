build:
	go run main.go

tests:
	go clean -testcache && cd ./test && go test *.go -v -cover

up:
	docker-compose up
	
down:
	docker-compose down