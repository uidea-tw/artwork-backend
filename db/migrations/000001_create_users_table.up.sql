CREATE TYPE gender AS ENUM ('F', 'M', 'O');

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "nickname" varchar,
    "introduction" varchar,
    "birth" timestamptz,
    "gender" gender NOT NULL,
    "username" varchar,
    "password" varchar,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("id");