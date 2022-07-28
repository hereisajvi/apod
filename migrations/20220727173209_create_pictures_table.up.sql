CREATE TABLE "pictures" (
    "id" UUID PRIMARY KEY NOT NULL UNIQUE,
    "copyright" VARCHAR NOT NULL,
    "date" TIMESTAMP  NOT NULL,
    "explanation" VARCHAR NOT NULL,
    "hdurl" VARCHAR NOT NULL,
    "media_type" VARCHAR NOT NULL,
    "service_version" VARCHAR NOT NULL,
    "title" VARCHAR NOT NULL,
    "url" VARCHAR NOT NULL UNIQUE
);
