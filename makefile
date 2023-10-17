# Migration Commands
c_m:
	@echo "Creating migrations..."
	migrate create -ext sql -dir db/migrations -seq $(name)

# PostgreSQL Commands
p_up:
	@echo "Starting PostgreSQL..."
	docker-compose up -d

p_down:
	@echo "Stopping and removing PostgreSQL..."
	docker-compose down

db_up:
	@echo "Creating a database..."
	docker exec -it finTech_postgres createdb --username=root --owner=root finTech_postgres_db
	docker exec -it finTech_postgres_live createdb --username=root --owner=root finTech_postgres_db_live

db_down:
	@echo "Dropping the database..."
	docker exec -it finTech_postgres dropdb --username=root finTech_postgres_db
	docker exec -it finTech_postgres_live dropdb --username=root finTech_postgres_db_live

# Migration Up and Down Commands
m_up:
	@echo "Applying database migrations (up)..."
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/finTech_postgres_db?sslmode=disable" up
	migrate -path db/migrations -database "postgres://root:secret@localhost:5433/finTech_postgres_db_live?sslmode=disable" up

m_down:
	@echo "Reverting database migrations (down)..."
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/finTech_postgres_db?sslmode=disable" down
	migrate -path db/migrations -database "postgres://root:secret@localhost:5433/finTech_postgres_db_live?sslmode=disable" down

# SQLC Commands
sqlc:
	@echo "Generating SQLC code..."
	sqlc generate

# Testing
test:
	@echo "Running tests..."
	go test -v -cover ./...

# package for JWT token
jwt:
	go get github.com/golang-jwt/jwt
# Start Development Server
start:
	@echo "Starting the development server..."
	CompileDaemon -command="./Fin-Remittance"
