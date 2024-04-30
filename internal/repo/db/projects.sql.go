// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: projects.sql

package dbrepo

import (
	"context"

	"github.com/hrz8/goatsapp/internal/repo/dto"
)

type CreateNewProjectsParams struct {
	Name        string              `db:"name" json:"name"`
	Alias       string              `db:"alias" json:"alias"`
	Description *string             `db:"description" json:"description"`
	Settings    dto.ProjectSettings `db:"settings" json:"settings"`
}

const getProjects = `-- name: GetProjects :many
SELECT id, name, alias, description, settings, created_at, updated_at FROM projects WHERE 1 = 1
`

func (q *Queries) GetProjects(ctx context.Context) ([]*Projects, error) {
	rows, err := q.db.Query(ctx, getProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Projects
	for rows.Next() {
		var i Projects
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Alias,
			&i.Description,
			&i.Settings,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
