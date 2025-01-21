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

migrate-up-online:
	migrate -path db/migrations -database "postgresql://root:7RaUKxW8F64IH5dmCODs2B1Z0tvj3rg9@free.clusters.zeabur.com:31414/zeabur?sslmode=disable" -verbose up

migrate-down-online:
	migrate -path db/migrations -database "postgresql://root:7RaUKxW8F64IH5dmCODs2B1Z0tvj3rg9@free.clusters.zeabur.com:31414/zeabur?sslmode=disable" -verbose down

.PHONY: migrate-up migrate-down