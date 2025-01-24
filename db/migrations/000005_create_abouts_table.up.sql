CREATE TABLE "abouts" (
    "id" bigserial PRIMARY KEY,
    "content" varchar NOT NULL,
    "cover" varchar NOT NULL,
    "cover_blur" varchar NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "abouts" ("id");