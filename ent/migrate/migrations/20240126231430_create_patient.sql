-- Create "patients" table
CREATE TABLE "patients" ("id" uuid NOT NULL, "name" character varying NOT NULL, "surname" character varying NOT NULL, "patronymic" character varying NOT NULL, "age" bigint NOT NULL, "gender" character varying NOT NULL, "country" character varying NOT NULL, PRIMARY KEY ("id"));
