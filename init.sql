CREATE TABLE IF NOT EXISTS "user" (
	"user_id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
	"username" VARCHAR (50) NOT NULL UNIQUE,
	"email" VARCHAR (50) NOT NULL UNIQUE,
	"password" VARCHAR (255) NOT NULL,
	"created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "category" (
    "category_id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "category_name" VARCHAR (255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "post" (
	"post_id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
	"title" VARCHAR (50) NOT NULL,
	"content" VARCHAR (255) NOT NULL,
	"user_id" SERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,
	"category_id" INT NOT NULL REFERENCES "category" ("category_id") ON DELETE CASCADE,
	"created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL	
);

CREATE TABLE IF NOT EXISTS "comment" (
    "comment_id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "content" VARCHAR (255) NOT NULL,
    "post_id" SERIAL NOT NULL REFERENCES "post" ("post_id") ON DELETE CASCADE,
    "user_id"  SERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,
    "created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "post_vote" (
    "post_vote_id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "value" BOOLEAN NOT NULL,
    "user_id" SERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,
    "post_id" SERIAL NOT NULL REFERENCES "post" ("post_id") ON DELETE CASCADE,
    "created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "comment_vote" (
    "comment_vote_id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "value" BOOLEAN NOT NULL,
    "user_id" SERIAL NOT NULL REFERENCES "user" ("user_id") ON DELETE CASCADE,    
    "comment_id" SERIAL NOT NULL REFERENCES "comment" ("comment_id") ON DELETE CASCADE,
    "created" TIMESTAMP NOT NULL,
	"updated" TIMESTAMP NOT NULL
);