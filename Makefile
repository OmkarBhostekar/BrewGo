new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateup:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/brewgo?sslmode=disable" -verbose up
migratedown:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/brewgo?sslmode=disable" -verbose down

mock-user-service:
	mockgen -destination ./services/user/db/mock/store.go github.com/omkarbhostekar/brewgo/services/user/db/sqlc Store

sqlc:
	sqlc generate

proto:
	rm -f proto/gen/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=proto/gen --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=brewgo\
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

run-api-gateway:
	cd services/api-gateway && go run main.go

run-user-service:
	cd services/user && go run main.go

run-product-service:
	cd services/product && go run main.go

run-order-service:
	cd services/order && go run main.go

.PHONY: new_migration migrateup migratedown proto mock-user-service run-user-service