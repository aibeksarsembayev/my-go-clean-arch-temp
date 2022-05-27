-- CREATE TABLE IF NOT EXISTS "post" (
-- 	"id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
-- 	"title" VARCHAR (50) NOT NULL,
-- 	"content" VARCHAR (255) NOT NULL,
-- 	"user" SERIAL NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
-- 	"category" INT NOT NULL,
-- 	"created" TIMESTAMP NOT NULL,
-- 	"updated" TIMESTAMP NOT NULL	
-- );

-- CREATE TABLE IF NOT EXISTS "user" (
-- 	"id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
-- 	"name" VARCHAR (50) NOT NULL UNIQUE,
-- 	"email" VARCHAR (50) NOT NULL UNIQUE,
-- 	"password" VARCHAR (255) NOT NULL,
-- 	"created" TIMESTAMP NOT NULL,
-- 	"updated" TIMESTAMP NOT NULL
-- );

CREATE TABLE IF NOT EXISTS "user" (
    "user_id" BIGSERIAL PRIMARY KEY,
    "email" TEXT NOT NULL UNIQUE,
    "username" VARCHAR(16) NOT NULL,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "post" (
    "post_id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGSERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,
    "title" VARCHAR(128) NOT NULL,
    "text" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "comment" (
    "comment_id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGSERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,
    "post_id" BIGSERIAL NOT NULL REFERENCES "post" ("post_id") ON DELETE CASCADE,
    "text" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "tag" (
    "tag_id" BIGSERIAL PRIMARY KEY,
    "name" TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "post_tag" (
    "post_tag_id" BIGSERIAL PRIMARY KEY,
    "post_id" BIGSERIAL NOT NULL REFERENCES "post" ("post_id") ON DELETE CASCADE,
    "tag_id" BIGSERIAL NOT NULL REFERENCES "tag" ("tag_id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "vote_post" (
    "vote_post_id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGSERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,
    "post_id" BIGSERIAL NOT NULL REFERENCES "post" ("post_id") ON DELETE CASCADE,
    "vote" SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS "vote_comment" (
    "vote_post_id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGSERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,
    "comment_id" BIGSERIAL NOT NULL REFERENCES "comment" ("comment_id") ON DELETE CASCADE,
    "vote" SMALLINT NOT NULL
);