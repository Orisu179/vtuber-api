
CREATE TABLE Platforms (
                           "platforms_id" SERIAL PRIMARY KEY,
                           "uri_template" VARCHAR NOT NULL
);

CREATE TABLE Groups (
                        "groups_id" SERIAL PRIMARY KEY,
                        "name" VARCHAR(30) NOT NULL,
                        "website" VARCHAR
);

CREATE TABLE Vtubers (
	"vtubers_id" UUID,
	"groups_id" SERIAL,
	"name_default" VARCHAR NOT NULL,
	"name_en" VARCHAR,
	"name_jp" VARCHAR,
	"name_cn" VARCHAR,
	"bio" TEXT,
	"languages" CHAR(2) NOT NULL,
	"debut_date" DATE, 
	PRIMARY KEY ("vtubers_id"),
	FOREIGN KEY ("groups_id") REFERENCES Groups("groups_id")
);
-- Relations
CREATE TABLE streamOn (
	"account_id" BIGSERIAL NOT NULL,
	"vtubers_id" BIGSERIAL NOT NULL,
	"platforms_id" SERIAL NOT NULL,
	PRIMARY KEY(vtubers_id, platforms_id),
	FOREIGN KEY(vtubers_id) REFERENCES Vtubers(vtubers_id),
	FOREIGN KEY(platforms_id) REFERENCES Platforms(platforms_id)
);
