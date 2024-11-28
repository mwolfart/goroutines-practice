download:
	go mod download

build:
	rm -rf bin
	go build -o ./bin/main .

run:
	go run .