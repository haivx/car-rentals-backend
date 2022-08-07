postgres:
		docker run --name postgres12  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
		docker exec -it postgres12 createdb --username=root --owner=root car_rental

dropdb:
		docker exec -it postgres12 dropdb --username=root --owner=root car_rental
	
migrateup:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/car_rental?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/car_rental?sslmode=disable" -verbose down

test:
		go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test