-- name: CreateGroup :one
INSERT INTO Groups (groups_id, name, website)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateAuthor :one
 INSERT INTO platforms(platforms_id, uri_template)
 VALUES ($1, $2)
 RETURNING *;