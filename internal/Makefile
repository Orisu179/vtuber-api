postgresCreate:
	docker run --name postgresAlpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgresAlpine createdb --username=root --owner=root vtuber

dropdb:
	docker exec -it postgresAlpine dropdb vtuber

migrateUp:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/vtuber?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/vtuber?sslmode=disable" -verbose down

.PHONY: postgresCreate createdb dropdb migrateUp migrateDown
