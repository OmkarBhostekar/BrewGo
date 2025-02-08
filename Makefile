new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateup:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/brewgo?sslmode=disable" -verbose up
migratedown:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/brewgo?sslmode=disable" -verbose down

.PHONY: new_migration migrateup migratedown