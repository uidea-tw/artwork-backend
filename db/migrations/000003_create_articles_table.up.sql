CREATE TABLE "articles" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "name" varchar NOT NULL,
    "desc" varchar NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    constraint fk_user_id foreign key ("user_id") references "users" ("id")
);

CREATE INDEX ON "articles" ("id");
