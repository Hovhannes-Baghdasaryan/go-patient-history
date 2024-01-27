ifneq (,$(wildcard ./local.env))
    include local.env
    export
    ENV_FILE_PARAM = --env-file local.env
endif

custom_migrate:
	atlas migrate new $(name) --dir file://ent/migrate/migrations

rehash:
	atlas migrate hash --dir file://ent/migrate/migrations

apply:
	atlas migrate apply --dir file://ent/migrate/migrations --url postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)

status:
	atlas migrate status --dir file://ent/migrate/migrations --url postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)

diff:
	atlas migrate diff $(name) --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?search_path=public&sslmode=disable"