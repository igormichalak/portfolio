include ../.env*

.PHONY: migrate/top
migrate/top:
	@tern migrate --migrations migrations --user $(POSTGRES_USER) --password $(POSTGRES_PASSWORD)

.PHONY: migrate/up
migrate/up:
	@tern migrate --destination +1 --migrations migrations --user $(POSTGRES_USER) --password $(POSTGRES_PASSWORD)

.PHONY: migrate/down
migrate/down:
	@tern migrate --destination -1 --migrations migrations --user $(POSTGRES_USER) --password $(POSTGRES_PASSWORD)
