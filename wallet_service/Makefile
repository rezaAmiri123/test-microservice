DB_URL=postgresql://postgres:postgres@api_postgesql:5432/wallet_service?sslmode=disable

createdb:
	docker exec -it api_postgesql createdb --username=postgres --owner=postgres wallet_service

dropdb:
	docker exec -it api_postgesql dropdb wallet_service
	
psql:
	psql -h api_postgesql -U postgres -d wallet_service -W

migratecreate:
	migrate create -ext sql -dir internal/adapters/migration -seq wallet_service

migrateup:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path internal/adapters/migration -database "$(DB_URL)" -verbose down 1

migrateforce:
	migrate  -path internal/adapters/migration -database "$(DB_URL)" force 1
