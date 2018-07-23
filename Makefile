prepare:
	protoc --go_out=plugins=grpc:. echo/*.proto

format:
	gofmt -s -w .

lint:
	gofmt -d .

run:
	make prepare
	go run main.go