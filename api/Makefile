-include ../.env*

.PHONY: default all

default:
	@echo "Please specify the target."

all: migrate-top migrate-up migrate-down

migrate-top:
	@tern migrate -m migrations --user ${POSTGRES_USER} --password ${POSTGRES_PASSWORD}

migrate-up:
	@tern migrate -d +1 -m migrations --user ${POSTGRES_USER} --password ${POSTGRES_PASSWORD}

migrate-down:
	@tern migrate -d -1 -m migrations --user ${POSTGRES_USER} --password ${POSTGRES_PASSWORD}
