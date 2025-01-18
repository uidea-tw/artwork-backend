create-migration: 
	@if [ -z "$(name)" ]; then \
		read -p "Enter migration name: " name; \
		migrate create -ext sql -dir db/migrations -seq $$name; \
	else \
		migrate create -ext sql -dir db/migrations -seq $(name); \
	fi
migrate-up:
	migrate -path db/migrations -database "postgresql://henry:password@localhost:5432/artwork?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://henry:password@localhost:5432/artwork?sslmode=disable" -verbose down

migrate-down-force:
	migrate -path db/migrations -database "postgresql://henry:password@localhost:5432/artwork?sslmode=disable" force 1

.PHONY: migrate-up migrate-down