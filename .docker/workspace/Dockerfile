FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

ENV PSQL_DBNAME=go-demo
ENV PSQL_HOST=postgresql
ENV PSQL_PORT=5432
ENV PSQL_USER=go-demo
ENV PSQL_PASS=password

# go command install
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
RUN go install github.com/rubenv/sql-migrate/...@v1.2.0