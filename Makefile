run:
	go run ./cmd/api

createMigrate:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrateUp:
	migrate -path=sql/migrations -database "postgres://postgres:postgres@localhost/sistema_de_passagem?sslmode=disable" -verbose up

migrateDown:
	migrate -path=sql/migrations -database "postgres://postgres:postgres@localhost/sistema_de_passagem?sslmode=disable" -verbose down