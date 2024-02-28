CREATE TABLE "vtubers" (
	"uuid" BIGSERIAL PRIMARY KEY AUTO_INCREMENT,
	"name_default" VARCHAR NOT NULL,
	"name_en" VARCHAR,
	"name_jp" VARCHAR,
	"name_cn" VARCHAR,
	"bio" TEXT,
--	"platforms" VARCHAR NOT NULL, 
	"youtube_link" VARCHAR NOT NULL,
	"languages" CHAR(2) NOT NULL,
	"debut_date" DATE, 
);

-- TODO Relate platforms to accounts instead of just youtube
-- CREATE TABLE "platforms" (
	-- "id" SERIAL PRIMARY KEY,
	-- "link_syntax" VARCHAR NOT NULL,
	-- "uri_template" VARCHAR NOT NULL,
-- );

CREATE TABLE "groups" (
	"id" SERIAL PRIMARY KEY,
	"vtuber_id" BIGSERIAL,
	"name" VARCHAR(30) NOT NULL,
	"website" VARCHAR,
);

ALTER TABLE "groups" ADD FOREIGN KEY ("vtuber_id") REFERENCES "vtuber" ("uuid");

CREATE INDEX ON "groups" ("vtuber_id")
