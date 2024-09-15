gen-migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5432/example?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:password@localhost:5432/example?sslmode=disable" down

seed:
	go run cmd/seed/main.go
