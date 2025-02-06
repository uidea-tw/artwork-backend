CREATE TABLE "files" (
    "id" char(36) PRIMARY KEY,
    "filename" varchar NOT NULL,
    "storage_name" varchar NOT NULL,
    "bucket" varchar(100) NOT NULL,
    "content_type" varchar NOT NULL,
    "size" int NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "files" ("id");