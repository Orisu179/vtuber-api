// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAuthor = `-- name: CreateAuthor :one
 INSERT INTO platforms(platforms_id, uri_template)
 VALUES ($1, $2)
 RETURNING platforms_id, name, uri_template
`

type CreateAuthorParams struct {
	PlatformsID int32
	UriTemplate string
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Platform, error) {
	row := q.db.QueryRow(ctx, createAuthor, arg.PlatformsID, arg.UriTemplate)
	var i Platform
	err := row.Scan(&i.PlatformsID, &i.Name, &i.UriTemplate)
	return i, err
}

const createGroup = `-- name: CreateGroup :one
INSERT INTO Groups (groups_id, name, website)
VALUES ($1, $2, $3)
RETURNING groups_id, name, website
`

type CreateGroupParams struct {
	GroupsID int32
	Name     string
	Website  pgtype.Text
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error) {
	row := q.db.QueryRow(ctx, createGroup, arg.GroupsID, arg.Name, arg.Website)
	var i Group
	err := row.Scan(&i.GroupsID, &i.Name, &i.Website)
	return i, err
}

const createLink = `-- name: CreateLink :one
INSERT INTO Links(account_id, vtubers_id, platforms_id)
VALUES ($1, $2, $3)
RETURNING account_id, vtubers_id, platforms_id
`

type CreateLinkParams struct {
	AccountID   int64
	VtubersID   pgtype.UUID
	PlatformsID int32
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (Link, error) {
	row := q.db.QueryRow(ctx, createLink, arg.AccountID, arg.VtubersID, arg.PlatformsID)
	var i Link
	err := row.Scan(&i.AccountID, &i.VtubersID, &i.PlatformsID)
	return i, err
}

const createVtuber = `-- name: CreateVtuber :one
INSERT INTO Vtubers(vtubers_id, groups_id, name_default, languages, debut_date)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING vtubers_id, groups_id, name_default, name_en, name_jp, name_cn, bio, languages, debut_date, gender
`

type CreateVtuberParams struct {
	GroupsID    pgtype.Int4
	NameDefault string
	Languages   []string
	DebutDate   pgtype.Date
}

func (q *Queries) CreateVtuber(ctx context.Context, arg CreateVtuberParams) (Vtuber, error) {
	row := q.db.QueryRow(ctx, createVtuber,
		arg.GroupsID,
		arg.NameDefault,
		arg.Languages,
		arg.DebutDate,
	)
	var i Vtuber
	err := row.Scan(
		&i.VtubersID,
		&i.GroupsID,
		&i.NameDefault,
		&i.NameEn,
		&i.NameJp,
		&i.NameCn,
		&i.Bio,
		&i.Languages,
		&i.DebutDate,
		&i.Gender,
	)
	return i, err
}

const getVtuber = `-- name: GetVtuber :one
SELECT vt.vtubers_id, vt.groups_id, vt.name_default, vt.name_en, vt.name_jp, vt.name_cn, vt.bio, vt.languages, vt.debut_date, vt.gender, g.name AS group_name
FROM Vtubers vt
         JOIN Groups g ON vt.groups_id = g.groups_id
WHERE vt.name_default LIKE $1 OR vt.name_en LIKE $1
   OR vt.name_jp LIKE $1 OR vt.name_cn LIKE $1
LIMIT 1
`

type GetVtuberRow struct {
	VtubersID   pgtype.UUID
	GroupsID    pgtype.Int4
	NameDefault string
	NameEn      pgtype.Text
	NameJp      pgtype.Text
	NameCn      pgtype.Text
	Bio         pgtype.Text
	Languages   []string
	DebutDate   pgtype.Date
	Gender      pgtype.Text
	GroupName   string
}

func (q *Queries) GetVtuber(ctx context.Context, nameDefault string) (GetVtuberRow, error) {
	row := q.db.QueryRow(ctx, getVtuber, nameDefault)
	var i GetVtuberRow
	err := row.Scan(
		&i.VtubersID,
		&i.GroupsID,
		&i.NameDefault,
		&i.NameEn,
		&i.NameJp,
		&i.NameCn,
		&i.Bio,
		&i.Languages,
		&i.DebutDate,
		&i.Gender,
		&i.GroupName,
	)
	return i, err
}

const getVtuberGroups = `-- name: GetVtuberGroups :many
SELECT vt.vtubers_id, vt.groups_id, vt.name_default, vt.name_en, vt.name_jp, vt.name_cn, vt.bio, vt.languages, vt.debut_date, vt.gender, g.name AS group_name
FROM Vtubers vt
JOIN Groups g ON vt.groups_id = g.groups_id
`

type GetVtuberGroupsRow struct {
	VtubersID   pgtype.UUID
	GroupsID    pgtype.Int4
	NameDefault string
	NameEn      pgtype.Text
	NameJp      pgtype.Text
	NameCn      pgtype.Text
	Bio         pgtype.Text
	Languages   []string
	DebutDate   pgtype.Date
	Gender      pgtype.Text
	GroupName   string
}

func (q *Queries) GetVtuberGroups(ctx context.Context) ([]GetVtuberGroupsRow, error) {
	rows, err := q.db.Query(ctx, getVtuberGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVtuberGroupsRow
	for rows.Next() {
		var i GetVtuberGroupsRow
		if err := rows.Scan(
			&i.VtubersID,
			&i.GroupsID,
			&i.NameDefault,
			&i.NameEn,
			&i.NameJp,
			&i.NameCn,
			&i.Bio,
			&i.Languages,
			&i.DebutDate,
			&i.Gender,
			&i.GroupName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVtuberPlatform = `-- name: GetVtuberPlatform :many
SELECT l.account_id, l.vtubers_id, l.platforms_id, p.uri_template, vt.name_default AS vtuber_name
FROM links l
    JOIN Platforms p ON l.platforms_id = p.platforms_id
    JOIN Vtubers vt ON l.vtubers_id = vt.vtubers_id
`

type GetVtuberPlatformRow struct {
	AccountID   int64
	VtubersID   pgtype.UUID
	PlatformsID int32
	UriTemplate string
	VtuberName  string
}

func (q *Queries) GetVtuberPlatform(ctx context.Context) ([]GetVtuberPlatformRow, error) {
	rows, err := q.db.Query(ctx, getVtuberPlatform)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVtuberPlatformRow
	for rows.Next() {
		var i GetVtuberPlatformRow
		if err := rows.Scan(
			&i.AccountID,
			&i.VtubersID,
			&i.PlatformsID,
			&i.UriTemplate,
			&i.VtuberName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
