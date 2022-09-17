backend:
	go run cmd/server/main.go

update_modules:
	go mod tidy

migrate_db:
	go run cmd/dbmigrate/main.go

start_api: update_modules migrate_db backend

start_fe:
	npm run dev
