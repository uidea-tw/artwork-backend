CREATE TABLE "admins" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "admins" ("id");