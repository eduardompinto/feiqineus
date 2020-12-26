# TODO make work with docker
# docker run --rm -v migrations:/migrations --network feiqineus_rede migrate/migrate -path=/migrations/  -database "postgres://postgres:password@database:5432/feiqineus?sslmode=disable" up
migrate -source file://migrations -database "postgres://postgres:password@localhost:5432/feiqineus?sslmode=disable" up