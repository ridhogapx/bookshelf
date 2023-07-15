postgres:
	docker run --name=postgres_container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

start-db:
	docker container start postgres_container

build-server:
	go build -o bin/server server/main.go

build-client:
	go build -o bin/client client/main.go

docker:
	docker build -t book-service
