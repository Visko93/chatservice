createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -database "postgresql://root:root@localhost:5432/chat_test?sslmode=disable" -path sql/migrations -verbose up

migratedown:
	migrate -database "postgresql://root:root@localhost:5432/chat_test?sslmode=disable" -path sql/migrations -verbose drop

grpc:
	protoc --go_out=. --go-grpc_out=. proto/chat.proto --experimental_allow_proto3_optional

.PHONY: migrate createmigration migratedown grpc