CREATE TABLE IF NOT EXISTS "post" (
	"id" int NOT NULL,
	"title" text NOT NULL,
	"user" int NOT NULL,
	"content" text NOT NULL,
	"category" int NOT NULL,
	"updated" time NOT NULL,
	"created" time NOT NULL,
	CONSTRAINT post_pk PRIMARY KEY ("id"),
	CONSTRAINT post_fk_1 FOREIGN KEY ("category") REFERENCES "category"("id"),
	CONSTRAINT post_fk FOREIGN KEY ("user") REFERENCES "user"("id")
);

CREATE TABLE IF NOT EXISTS "category" (
	"id" int NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT category_pk PRIMARY KEY ("id"),
	CONSTRAINT category_un UNIQUE ("id"),
	CONSTRAINT category_un UNIQUE ("name")
);

CREATE TABLE IF NOT EXISTS "comment" (
	"id" int NOT NULL,
	"content" text NOT NULL,
	"post" int NOT NULL,
	"user" int NOT NULL,
	"created" time NOT NULL,
	"updated" time NOT NULL,
	CONSTRAINT comment_pk PRIMARY KEY ("id"),
	CONSTRAINT comment_un UNIQUE ("id"),
	CONSTRAINT comment_fk FOREIGN KEY ("post") REFERENCES "post"("id"),
	CONSTRAINT comment_fk_1 FOREIGN KEY ("user") REFERENCES "user"("id")
);

CREATE TABLE IF NOT EXISTS "user" (
	"id" int NOT NULL,
	"password" text NOT NULL,
	"email" text NOT NULL,
	"created" time NOT NULL,
	"updated" time NOT NULL,
	"name" text NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY ("id"),
	CONSTRAINT user_un UNIQUE ("id","password","email")
);
