.PHONY: migrateup migratedown createMigration run
migrateup:
	migrate -path backend/pkg/db/migrations/sqlite -database sqlite://./backend/social_network.db -verbose up
migratedown:
	migrate -path backend/pkg/db/migrations/sqlite -database sqlite://./backend/social_network.db -verbose down
createMigration:
	migrate create -ext sql -dir backend/pkg/db/migrations/sqlite -seq create_social_network_schema
runGo:
	rm -rf backend/social_network.db
	cd backend && go run populate/main.go
	cd backend && go run server.go
runJs:
	cd frontend && npm run dev
runTest:
	rm -rf backend/Test/social_network_test.db
	cd backend/Test && go test -v .


backendDocker:
	docker build -t golang-server ./backend
	docker run --rm -it  -p 5070:5070/tcp golang-server:latest

frontendDocker:
	docker build -t react-frontend-app ./frontend
	docker run --rm -it -p 80:80/tcp -p 8070:8070/tcp react-frontend-app:latest