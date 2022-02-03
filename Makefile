.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti --network host adminer

platform:
	export PGUSER="postgres" && \
	export PGHOST="localhost" && \
	export PGDATABASE="postgres" && \
	export PGPASSWORD="secret" && \
	reflex -s go run cmd/server/main.go
