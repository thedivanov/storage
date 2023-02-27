build:
	docker build .
up:
	docker-compose up
vendor:
	cd src \
	go mod vendor
tidy:
	cd src \
	go mod tidy

protoc:
	protoc --go_out=./src --go_opt=paths=source_relative \
    --go-grpc_out=./src --go-grpc_opt=paths=source_relative \
    ./proto/*.proto