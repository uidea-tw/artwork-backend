CREATE TABLE "abouts" (
    "id" bigserial PRIMARY KEY,
    "content" varchar NOT NULL,
    "cover" varchar NOT NULL,
    "cover_blur" varchar NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    constraint fk_cover foreign key ("cover") references "files" ("id"),
    constraint fk_cover_blur foreign key ("cover_blur") references "files" ("id")
);

CREATE INDEX ON "abouts" ("id");