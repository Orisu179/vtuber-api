CREATE TABLE Vtubers (
	"uuid" BIGSERIAL AUTO_INCREMENT,
	"groups_id" SERIAL,
	"name_default" VARCHAR NOT NULL,
	"name_en" VARCHAR,
	"name_jp" VARCHAR,
	"name_cn" VARCHAR,
	"bio" TEXT,
--	"platforms" VARCHAR NOT NULL, 
	"youtube_link" VARCHAR NOT NULL,
	"languages" CHAR(2) NOT NULL,
	"debut_date" DATE, 
	PRIMARY KEY ("uuid"),
	FOREIGN KEY ("groups_id") REFERENCES Groups("id")
);

-- TODO Relate platforms to accounts instead of just youtube
-- CREATE TABLE "platforms" (
	-- "id" SERIAL PRIMARY KEY,
	-- "link_syntax" VARCHAR NOT NULL,
	-- "uri_template" VARCHAR NOT NULL,
-- );

CREATE TABLE Groups (
	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(30) NOT NULL,
	"website" VARCHAR,
);
