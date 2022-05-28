CREATE TABLE IF NOT EXISTS "user" (
	"id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
	"name" VARCHAR (50) NOT NULL UNIQUE,
	"email" VARCHAR (50) NOT NULL UNIQUE,
	"password" VARCHAR (255) NOT NULL,
	"created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "category" (
    "id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "name" VARCHAR (255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "post" (
	"id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
	"title" VARCHAR (50) NOT NULL,
	"content" VARCHAR (255) NOT NULL,
	"user_id" SERIAL NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
	"category_id" INT NOT NULL REFERENCES "category" ("id") ON DELETE CASCADE,
	"created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL	
);

CREATE TABLE IF NOT EXISTS "comment" (
    "id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "content" VARCHAR (255) NOT NULL,
    "post_id" SERIAL NOT NULL REFERENCES "post" ("id") ON DELETE CASCADE,
    "user_id"  SERIAL NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    "created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "post_vote" {
    "id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "value" BOOLEAN NOT NULL,
    "user_id" SERIAL NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    "post_id" SERIAL NOT NULL REFERENCES "post" ("id") ON DELETE CASCADE,
    "created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
};

CREATE TABLE IF NOT EXISTS "comment_vote" {
    "id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "value" BOOLEAN NOT NULL,
    "user_id" SERIAL NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,    
    "comment_id" SERIAL NOT NULL REFERENCES "comment" ("id") ON DELETE CASCADE,
    "created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
};