createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -database "postgresql://root:root@localhost:5432/chat_test?sslmode=disable" -path sql/migrations -verbose up

migratedown:
	migrate -database "postgresql://root:root@localhost:5432/chat_test?sslmode=disable" -path sql/migrations -verbose drop

.PHONY: migrate migratedown createmigration