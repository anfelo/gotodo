.PHONY: postgres adminer migrate platform

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti --network host adminer

platform:
	export DB_USERNAME="postgres" && \
	export DB_HOST="localhost" && \
	export DB_TABLE="postgres" && \
	export DB_PASSWORD="secret" && \
	export DB_PORT="5432" && \
	export SSL_MODE="disable" && \
	reflex -s go run cmd/server/main.go
