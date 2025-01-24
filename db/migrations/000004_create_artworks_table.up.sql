CREATE TABLE "artworks" (
    "id" bigserial PRIMARY KEY,
    "title" varchar(30) NOT NULL,
    "cover" varchar NOT NULL,
    "description" varchar NOT NULL,
    "content" varchar NOT NULL,
    "author"  varchar NOT NULL,
    "year" integer NOT NULL,
    "length" integer NOT NULL,
    "width" integer NOT NULL,
    "height" integer NOT NULL,
    "price" integer NOT NULL,
    "cover_blur" varchar NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "artworks" ("id");