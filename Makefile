postgres:
	docker run --name=postgres_container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

start-db:
	docker container start postgres_container

