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
	atlas migrate apply --dir file://ent/migrate/migrations --url postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?sslmode=disable

status:
	atlas migrate status --dir file://ent/migrate/migrations --url postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?sslmode=disable

diff:
	atlas migrate diff $(name) --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "docker://postgres/15/dev?search_path=public"
