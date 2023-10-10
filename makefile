c_m:
	#create migrations
	migrate create -ext sql -dir db/migrations -seq $(name)
p_up:
	# postgres up - create postgres server
	docker-compose up -d
p_down:
	# postgres down - delete postgres server
	docker-compose down
db_up:
	docker exec -it finTech_postgres createdb --username=root --owner=root finTech_postgres_db

db_down:
	docker exec -it finTech_postgres dropdb --username=root finTech_postgres_db
m_up:
	#run migrate up
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/finTech_postgres_db?sslmode=disable" up
m_down:
	#run migrate down
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/finTech_postgres_db?sslmode=disable" down

sqlc:
	sqlc generate

start:
	CompileDaemon -command="./Fin-Remittance"







    