-- name: CreateGroup :one
INSERT INTO Groups (groups_id, name, website)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateAuthor :one
 INSERT INTO platforms(platforms_id, uri_template)
 VALUES ($1, $2)
 RETURNING *;

-- name: CreateVtuber :one
INSERT INTO Vtubers(vtubers_id, groups_id, name_default, languages, debut_date)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING *;

-- name: CreateLink :one
INSERT INTO Links(account_id, vtubers_id, platforms_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetVtuberGroups :many
SELECT vt.*, g.name AS group_name
FROM Vtubers vt
JOIN Groups g ON vt.groups_id = g.groups_id;

-- name: GetVtuberPlatform :many
SELECT l.*, p.uri_template, vt.name_default AS vtuber_name
FROM links l
    JOIN Platforms p ON l.platforms_id = p.platforms_id
    JOIN Vtubers vt ON l.vtubers_id = vt.vtubers_id;

-- name: GetVtuber :one
SELECT vt.*, g.name AS group_name
FROM Vtubers vt
         JOIN Groups g ON vt.groups_id = g.groups_id
WHERE vt.name_default LIKE $1 OR vt.name_en LIKE $1
   OR vt.name_jp LIKE $1 OR vt.name_cn LIKE $1
LIMIT 1;
