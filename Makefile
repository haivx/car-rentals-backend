postgres:
		docker run --name postgres12  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
		docker exec -it postgres12 createdb --username=root --owner=root car_rental

dropdb:
		docker exec -it postgres12 dropdb --username=root --owner=root car_rental
	
migrateup:
		migrate -path db/migration -database "postgresql://root:0yU0uRlWY6RNLj9vukvP@rental-car.cjvnnidpz82k.us-east-1.rds.amazonaws.com:5432/rentalcar?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/car_rental?sslmode=disable" -verbose down

test:
		go test -v -cover ./...

proto:
	rm -f pb/*.go
		protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
				--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
				proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown test proto